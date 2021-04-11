package main

import (
	"fmt"
	"study-go/EasyGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}