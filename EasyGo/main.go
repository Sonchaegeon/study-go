package main

import (
	"fmt"
	"study-go/EasyGo/accounts"
)

func main() {
	account := accounts.NewAccount("sonchaegeon")
	fmt.Println(account)
}