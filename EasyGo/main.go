package main

import "fmt"

func main() {
	const name string = "손채건"
	// name = "qweqwe" -> error!
	var age int = 17
	autoType := false
	age++
	fmt.Println(age)
	fmt.Println(autoType)
}
