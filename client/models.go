package client

type BurnPayload struct {
	TxnID           string `json:"0chain_txn_id"`
	Nonce           int64  `json:"nonce"`
	Amount          int64  `json:"amount"`
	EthereumAddress string `json:"ethereum_address"`
}

type AuthorizerNode struct {
	PublicKey string `json:"public_key"`
	URL       string `json:"url"`
}