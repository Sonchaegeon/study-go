package main

import (
	"fmt"
	"study-go/EasyGo/accounts"
)

func main() {
	account := accounts.NewAccount("sonchaegeon")
	account.Deposit(200000)
	fmt.Println(account)
}