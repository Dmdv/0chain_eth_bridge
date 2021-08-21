package client

import (
	"encoding/json"
	"fmt"
	"github.com/0chain/gosdk/zcncore"
)

func TestStorageSc() zcncore.TransactionScheme {
	fmt.Println("----------------------------------------------")
	fmt.Println("Started executing smart contract TestStorageSc...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}


	status.Begin()
	err = txn.ExecuteSmartContract(StorageAddress, "state_error_test", "", zcncore.ConvertToValue(1))
	if err != nil {
		fmt.Printf("Transaction failed with error: '%s'", err.Error())
		return nil
	}

	status.Wait()
	fmt.Printf("Executed smart contract TestStorageSc with TX = '%s'\n", txn.GetTransactionHash())

	VerifyTransaction(txn, status)

	return txn
}

func TestZcnscSc() zcncore.TransactionScheme {
	fmt.Println("----------------------------------------------")
	fmt.Println("Started executing smart contract TestZcnscSc...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}


	status.Begin()
	err = txn.ExecuteSmartContract(ZcnscAddress, "state_error_test", "", zcncore.ConvertToValue(1))
	if err != nil {
		fmt.Printf("Transaction failed with error: '%s'", err.Error())
		return nil
	}

	status.Wait()
	fmt.Printf("Executed smart contract TestZcnscSc with TX = '%s'\n", txn.GetTransactionHash())

	VerifyTransaction(txn, status)

	return txn
}

func TestZcnscSc2() zcncore.TransactionScheme {
	fmt.Println("----------------------------------------------")
	fmt.Println("Started executing smart contract TestZcnscSc2...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}

	payload := &AuthorizerNode{
		PublicKey: "PublicKey",
		URL:       "localhost",
	}

	buffer, _ := json.Marshal(payload)

	status.Begin()
	err = txn.ExecuteSmartContract(ZcnscAddress, "state_error_test2", string(buffer), zcncore.ConvertToValue(1))
	if err != nil {
		fmt.Printf("Transaction failed with error: '%s'", err.Error())
		return nil
	}

	status.Wait()
	fmt.Printf("Executed smart contract TestZcnscSc2 with TX = '%s'\n", txn.GetTransactionHash())

	VerifyTransaction(txn, status)

	return txn
}

func TestDeleteAuthorizer() zcncore.TransactionScheme {
	fmt.Println("----------------------------------------------")
	fmt.Println("Started executing smart contract TestDeleteAuthorizer...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}

	status.Begin()
	err = txn.ExecuteSmartContract(ZcnscAddress, "state_error_test_delete", "", zcncore.ConvertToValue(1))
	if err != nil {
		fmt.Printf("Transaction failed with error: '%s'", err.Error())
		return nil
	}

	status.Wait()
	fmt.Printf("Executed smart contract TestDeleteAuthorizer with TX = '%s'\n", txn.GetTransactionHash())

	VerifyTransaction(txn, status)

	return txn
}
