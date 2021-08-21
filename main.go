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

	client.TestStorageSc()
	client.TestZcnscSc()
	client.TestDeleteAuthorizer()
	client.TestZcnscSc2()

	//client.PourTokens(100)
	//client.PourTokens(100)
	//client.PourTokens(100)
	//client.PourTokens(100)
	//client.CheckBalance()

	// Burn-tokens case
	// Description:
	// 1. User burns token
	// 2. Authorizer sends the client proof-of-Burn ticket
	// 3. User gathers tickets from authorizers
	//client.Burn(2, 4)
	//client.CheckBalance()

	// Add Authorizer
	// Description:
	// Authorizer is being registered only with PublicKey.
	// To check `AddAuthorizer` function we only need to call this function and verify transaction
	//client.RegisterAuthorizer("http://localhost:9999")

	// Mint tokens -> Check transaction
}