package main

import (
	"fmt"

	"github.com/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("min")
	fmt.Println(account)
}
