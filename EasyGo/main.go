package main

import (
	"fmt"
	"log"
	"study-go/EasyGo/accounts"
)

func main() {
	account := accounts.NewAccount("sonchaegeon")
	account.Deposit(200000)
	fmt.Println(account.Balance())
	err := account.Withdraw(300000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}