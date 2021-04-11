package main

import "fmt"

func main() {
	scg := map[string]string{"name": "sonchaegeon", "age": "18"}
	fmt.Println("name: " + scg["name"] + " age: " + scg["age"])

	for _, value := range scg {
		fmt.Println(value)
	}
}
