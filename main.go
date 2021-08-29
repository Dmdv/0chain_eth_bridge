package main

import (
	"eth_bridge/client"
	"eth_bridge/usecases"
	"fmt"
)

func main() {
	fmt.Println("Started e2e testing")

	// Preparations

	client.InitClient()
	client.CheckBalance()
	client.PourTokens(100)
	client.CheckBalance()

	// Converting from ZCN to WZCN

	for i := 1; i < 4; i++ {
		usecases.ToWzcn(1, int64(i))
	}
}