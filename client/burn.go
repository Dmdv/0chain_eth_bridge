package client

import (
	"encoding/json"
	"fmt"
	"github.com/0chain/gosdk/zcncore"
)

func Burn(amount float64, nonce int64) zcncore.TransactionScheme {
	fmt.Println("----------------------------------------------")
	fmt.Println("Started executing smart contract ZCNSC:Burn...")
	status := NewZCNStatus()
	txn, err := zcncore.NewTransaction(status, 0)
	if err != nil {
		ExitWithError(err)
	}

	payload := BurnPayload{
		TxnID:           "",
		Nonce:           nonce,
		Amount:          zcncore.ConvertToValue(amount),
		EthereumAddress: "ABC",
	}

	buffer, _ := json.Marshal(payload)

	fmt.Printf("Payload: BurnPayload: %s", buffer)

	status.Begin()
	err = txn.ExecuteSmartContract(
		ZcnscAddress,
		BurnMethod,
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
