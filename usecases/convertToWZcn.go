package usecases

import (
	"eth_bridge/authorizers"
	"eth_bridge/client"
)

// ToWzcn converts from ZCN to WZNC
// Description:
// 1. User burns token
// 2. Authorizer sends the client proof-of-Burn ticket
// 3. User gathers tickets from authorizers
// 4. User sends tickets to ETH bridge
func ToWzcn(amount float64, nonce int64) {
	client.Burn(amount, nonce)
	client.CheckBalance()

	authorizersFromChain := authorizers.GetAuthorizersFromChain()

	client.GetBurnProofTickets(authorizersFromChain)
	client.SendTicketsToEthereumBridge()
}
