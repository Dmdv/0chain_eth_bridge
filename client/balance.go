package client

import (
	"fmt"
	"github.com/0chain/gosdk/zcncore"
)

func CheckBalance() {
	fmt.Println("Started Checking balance...")
	balance := NewZCNStatus()
	balance.Begin()
	err := zcncore.GetBalance(balance)
	if err == nil {
		balance.Wait()
		fmt.Printf("Client balance: %d", balance.balance)
	} else {
		fmt.Println("Failed to get the balance: " + err.Error())
	}
}
