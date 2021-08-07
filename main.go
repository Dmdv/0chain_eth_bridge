package main

import (
	"fmt"

	"eth_bridge/code"
	"github.com/0chain/gosdk/zcncore"
	// "github.com/spf13/cobra"
)

func main(){
	if err := code.MakeConfig(); err != nil {
		code.ExitWithError(err)
	}

	statusBar := code.NewZCNStatus()
	txn, err := zcncore.NewTransaction(statusBar, 0)
	if err != nil {
		code.ExitWithError(err)
	}

	statusBar.Begin()
	err = txn.ExecuteSmartContract(
		zcncore.FaucetSmartContractAddress,
		"pour",
		"new wallet",
		zcncore.ConvertToValue(10),
	)
	if err == nil {
		statusBar.Wait()
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println("Executed transaction: " + txn.GetTransactionHash())
	fmt.Println("Verifying...")
	err = txn.Verify()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt. Println(txn.GetVerifyOutput())
}
