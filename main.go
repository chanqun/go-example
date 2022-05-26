package main

import (
	"fmt"
	"github.com/chanqun/go-example/accounts"
)

func main() {
	account := accounts.NewAccount("chanqun")

	fmt.Println(account)
}
