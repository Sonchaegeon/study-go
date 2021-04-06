package main

import (
	"fmt"
	"study-go/EasyGo/something"
)

func main() {
	fmt.Println("hello world")
	something.SayHello()
	something.sayBye() // error
}
