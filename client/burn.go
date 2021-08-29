package client

import (
	"encoding/json"
	"github.com/0chain/gosdk/zcncore"
)

func Burn(amount float64, nonce int64) zcncore.TransactionScheme {
	payload := BurnPayload {
		TxnID:           "",
		Nonce:           nonce,
		Amount:          zcncore.ConvertToValue(amount),
		EthereumAddress: "ABC",
	}

	buffer, _ := json.Marshal(payload)

	return StartAndVerifyTransaction(
		"ZCNSC",
		"BurnMethod",
		ZcnscAddress,
		string(buffer),
		amount,
	)
}
