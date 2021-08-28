package main

import (
	"eth_bridge/client"
	"fmt"
	// "github.com/spf13/cobra"
)

func main() {
	fmt.Println("Started e2e testing")

	// Preparations

	client.InitClient()
	client.CheckBalance()
	client.PourTokens(100)

	// Burn-tokens case
	// Description:
	// 1. User burns token
	// 2. Authorizer sends the client proof-of-Burn ticket
	// 3. User gathers tickets from authorizers
	// 4. User sends tickets to ETH bridge
	client.Burn(2, 4)
	client.CheckBalance()

	// Add Authorizer (not working yet)
	// Description:
	// To check `AddAuthorizer` function we only need to call this function and verify transaction
	// client.RegisterAuthorizer("http://localhost:9999")

	// Mint-tokens case

}