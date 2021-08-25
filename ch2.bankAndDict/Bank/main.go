package main

import (
	"fmt"

	"github.com/meanhide/nomadcoder/ch2.bankAndDict/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	// 	err := account.Withdraw(20)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
