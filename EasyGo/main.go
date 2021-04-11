package main

import "fmt"

func main() {
	names := []string{"kim", "lee", "son"}
	names = append(names, "park")
	fmt.Println(names)
}
