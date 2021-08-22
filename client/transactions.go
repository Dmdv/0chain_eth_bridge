package client

import (
	"fmt"
	"github.com/0chain/gosdk/zcncore"
)

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
