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
	// Burn tokens -> Check transaction
	// Mint tokens -> Check transaction
}