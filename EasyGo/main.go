package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)
	a = 123
	fmt.Println(*b)
}
