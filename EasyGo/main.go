package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood := []string{"pizza", "hamburger"}
	scg := person{"sonchaegeon", 18, favFood}
	fmt.Println(scg)
	fmt.Println(scg.favFood)
}
