package client

import (
	"encoding/json"
	"fmt"
	"github.com/0chain/gosdk/zcncore"
)

const (
	// ADDRESS ...
	ADDRESS = "6dba10422e368813802877a85039d3985d96760ed844092319743fb3a76712e0"
	name    = "zcn"
)

func GetAuthorizers() {
}

func AddAuthorizer() {
}

func Burn(amount float64, nonce int64) zcncore.TransactionScheme {
	fmt.Println("----------------------------------------------")
	fmt.Println("Started executing smart contract ZCNSC:Burn...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}

	payload := BurnPayload{
		Nonce:           nonce,
		EthereumAddress: "ABC",
	}

	buffer, _ := json.Marshal(payload)

	fmt.Printf("Payload: BurnPayload: %s", buffer)

	status.Begin()
	err = txn.ExecuteSmartContract(
		ADDRESS,
		"burn",
		string(buffer),
		zcncore.ConvertToValue(amount),
	)

	if err != nil {
		fmt.Printf("Transaction failed with error: '%s'", err.Error())
		return nil
	}

	status.Wait()
	fmt.Printf("Executed smart contract ZCNSC:Burn with TX = '%s'\n", txn.GetTransactionHash())

	VerifyTransaction(txn, status)

	return txn
}

func PourTokens(amount float64) zcncore.TransactionScheme {
	fmt.Println("----------------------------------------------")
	fmt.Println("Started executing smart contract Faucet:Pour...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}

	status.Begin()
	err = txn.ExecuteSmartContract(
		zcncore.FaucetSmartContractAddress,
		"pour",
		"new wallet",
		zcncore.ConvertToValue(amount),
	)

	if err != nil {
		fmt.Printf("Transaction failed with error: '%s'", err.Error())
		return nil
	}

	status.Wait()
	fmt.Printf("Executed smart contract Faucet:Pour with TX = '%s'\n", txn.GetTransactionHash())

	VerifyTransaction(txn, status)

	return txn
}

func VerifyTransaction(txn zcncore.TransactionScheme, status *ZCNStatus) {
	if txn == nil {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Printf("Started Verifying Transaction '%s' ...\n", txn.GetTransactionHash())

	status.Begin()
	err := txn.Verify()
	if err == nil {
		status.Wait()
		fmt.Printf("Verify output: '%s'\n", txn.GetVerifyOutput())
		fmt.Println("Verification completed OK")
	} else {
		fmt.Printf("Failed to verify with error: %s\n", err.Error())
		fmt.Printf("Verify error: '%s'\n" + txn.GetVerifyError())
		fmt.Printf("Transaction error: '%s'\n" + txn.GetTransactionError())
		fmt.Println("Verification FAILED")
	}
}
