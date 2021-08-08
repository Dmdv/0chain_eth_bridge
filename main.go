package main

import (
	"eth_bridge/client"
	"fmt"
	"github.com/0chain/gosdk/zcncore"
	// "github.com/spf13/cobra"
)

func main() {
	fmt.Println("Started e2e testing")
	client.InitClient()

	txn, status := PourTokens(10)
	VerifyTransactionWithDiagnostics(txn, status)

	client.CheckBalance()

	// Burn tokens -> Check transaction
	// Mint tokens -> Check transaction
	// Add Authorizer
	// Get list of authorizers
}

func PourTokens(amount float64) (zcncore.TransactionScheme, *client.ZCNStatus) {
	status := client.NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		client.ExitWithError(err)
	}

	status.Begin()
	err = txn.ExecuteSmartContract(
		zcncore.FaucetSmartContractAddress,
		"pour",
		"new wallet",
		zcncore.ConvertToValue(amount),
	)

	if err == nil {
		status.Wait()
		fmt.Printf("Executed smart contract Faucet:Pour with TX = '%s'\n", txn.GetTransactionHash())
		return txn, status
	}

	fmt.Printf("Transaction failed with error: '%s'", err.Error())
	return nil, status
}

func VerifyTransactionWithDiagnostics(txn zcncore.TransactionScheme, status *client.ZCNStatus) {
	if txn == nil {
		return
	}

	fmt.Printf("Verifying transaction '%s' ...\n", txn.GetTransactionHash())

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
