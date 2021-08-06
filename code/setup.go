package code

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/0chain/gosdk/core/zcncrypto"
	"github.com/0chain/gosdk/zcncore"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

const (
	configFile = "0chain.yaml"
)

var cfgFile string
var walletFile string
var dir string
var silent bool

var clientConfig string
var minSubmit int
var minCfm int
var CfmChainLength int

var clientWallet *zcncrypto.Wallet

func init() {
	cfgFileFlag := flag.String("config", "", "config file (default is 0chain.yaml)")
	walletFileFlag := flag.String("wallet", "", "wallet file (default is wallet.json)")
	dirFlag := flag.String("configDir", "", "configuration directory (default is $HOME/.zcn)")
	silentFlag := flag.Bool("silent", false, "Do not print sdk logs in stderr (prints by default)")

	flag.Parse()

	cfgFile = *cfgFileFlag
	walletFile = *walletFileFlag
	dir = *dirFlag
	silent = *silentFlag
}

func MakeConfig() {
	fmt.Println("Started e2e testing")
	initConfig()
}

func initConfig() {
	chainConfig := viper.New()

	var configDir string
	if dir != "" {
		configDir = dir
	} else {
		configDir = getConfigDir()
	}
	chainConfig.AddConfigPath(configDir)
	if &cfgFile != nil && len(cfgFile) > 0 {
		chainConfig.SetConfigFile(configDir + "/" + cfgFile)
	} else {
		chainConfig.SetConfigFile(configDir + "/" + configFile)
	}

	if err := chainConfig.ReadInConfig(); err != nil {
		ExitWithError("Can't read config:", err)
	}

	blockWorker := chainConfig.GetString("block_worker")
	signScheme := chainConfig.GetString("server_chain.signature_scheme")
	chainID := chainConfig.GetString("server_chain.id")
	minSubmit = chainConfig.GetInt("server_chain.min_submit")
	minCfm = chainConfig.GetInt("server_chain.min_confirmation")
	CfmChainLength = chainConfig.GetInt("server_chain.confirmation_chain_length")

	var walletFilePath string
	if &walletFile != nil && len(walletFile) > 0 {
		walletFilePath = path.Join(configDir, walletFile)
	} else {
		walletFilePath = path.Join(configDir, "wallet.json")
	}

	zcncore.SetLogFile("cmdlog.log", !silent)

	err := zcncore.InitZCNSDK(
		blockWorker,
		signScheme,
		zcncore.WithChainID(chainID),
		zcncore.WithMinSubmit(minSubmit),
		zcncore.WithMinConfirmation(minCfm),
		zcncore.WithConfirmationChainLength(CfmChainLength),
	)
	if err != nil {
		ExitWithError(err.Error())
	}

	var fresh bool

	if _, err := os.Stat(walletFilePath); os.IsNotExist(err) {
		fmt.Println("No wallet in path ", walletFilePath, "found. Creating wallet...")
		wg := &sync.WaitGroup{}
		statusBar := &ZCNStatus{wg: wg}

		wg.Add(1)
		err = zcncore.CreateWallet(statusBar)
		if err == nil {
			wg.Wait()
		} else {
			ExitWithError(err.Error())
		}

		if len(statusBar.walletString) == 0 || !statusBar.success {
			ExitWithError("Error creating the wallet." + statusBar.errMsg)
		}
		fmt.Println("ZCN wallet created!!")
		clientConfig = statusBar.walletString
		fmt.Println("Wallet string: " + clientConfig)
		file, err := os.Create(walletFilePath)
		if err != nil {
			ExitWithError(err.Error())
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		_, _ = fmt.Fprintf(file, clientConfig)

		fresh = true

	} else {
		f, err := os.Open(walletFilePath)
		if err != nil {
			ExitWithError("Error opening the wallet", err)
		}
		clientBytes, err := ioutil.ReadAll(f)
		if err != nil {
			ExitWithError("Error reading the wallet", err)
		}
		clientConfig = string(clientBytes)
	}

	wallet := &zcncrypto.Wallet{}
	err = json.Unmarshal([]byte(clientConfig), wallet)
	clientWallet = wallet
	if err != nil {
		ExitWithError("Invalid wallet at path:" + walletFilePath)
	}
	wg := &sync.WaitGroup{}
	err = zcncore.SetWalletInfo(clientConfig, false)
	if err == nil {
		wg.Wait()
	} else {
		ExitWithError(err.Error())
	}

	if fresh {
		log.Print("Creating related read pool for storage of smart-contract...")
		if err = createReadPool(); err != nil {
			log.Fatalf("Failed to create read pool: %v", err)
		}
		log.Printf("Read pool created successfully")
	}

	wg = &sync.WaitGroup{}
	statusBar := &ZCNStatus{wg: wg}
	wg.Add(1)
	_ = zcncore.RegisterToMiners(clientWallet, statusBar)
	wg.Wait()
	if statusBar.success {
		fmt.Println("Wallet registered at miners: ")
		fmt.Println("Wallet ClientID: " + clientWallet.ClientID)
		fmt.Println("Wallet ClientKey: " + clientWallet.ClientKey)
	} else {
		PrintError("Wallet registration failed. " + statusBar.errMsg)
		os.Exit(1)
	}
}

func getConfigDir() string {
	if dir != "" {
		return dir
	}
	var configDir string
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configDir = path.Join(home, ".zcn")
	return configDir
}
