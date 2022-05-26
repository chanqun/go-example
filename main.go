package main

import (
	"fmt"
	"github.com/chanqun/go-example/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("chanqun")

	account.Deposit(100)

	err := account.Withdraw(200)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance())
}
