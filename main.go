package main

import (
	"eth_bridge/client"
	"fmt"
	// "github.com/spf13/cobra"
)

func main() {
	fmt.Println("Started e2e testing")
	client.InitClient()
	client.PourTokens(10)
	client.CheckBalance()

	// Add Authorizer
	// Get list of authorizers

	// Burn tokens case -> Check transaction
	// 1. User burns token
	// 2. Authorizer signs the burn ticket
	// 3. User gets the ticket and sends to the wzcn to mint tokens

	client.Burn(2, 2)
	client.CheckBalance()

	// TODO: compare balances
	// TODO: Run multiple threads

	// Mint tokens -> Check transaction
}