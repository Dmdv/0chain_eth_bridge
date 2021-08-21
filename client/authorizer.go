package client

import (
	"encoding/json"
	"fmt"
	"github.com/0chain/gosdk/zcncore"
)

func RegisterAuthorizer(url string) zcncore.TransactionScheme {
	fmt.Println("---------------------------")
	fmt.Println("Started Registering an authorizer...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}

	payload := &AuthorizerNode{
		PublicKey: "public key",
		URL:       url,
	}

	buffer, _ := json.Marshal(payload)

	fmt.Printf("Payload: AuthorizerNode: %s\n", buffer)

	status.Begin()
	err = txn.ExecuteSmartContract(
		ZcnscAddress,
		AddAuthorizerMethod,
		string(buffer),
		zcncore.ConvertToValue(3),
	)

	if err != nil {
		fmt.Printf("Transaction failed with error: '%s'\n", err.Error())
		return nil
	}

	status.Wait()
	fmt.Printf("Executed smart contract ZCNSC:AddAuthorizer with TX = '%s'\n", txn.GetTransactionHash())

	VerifyTransaction(txn, status)

	return txn
}