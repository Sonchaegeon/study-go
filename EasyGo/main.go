package main

import (
	"fmt"
	"study-go/EasyGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	definition, _ := dictionary.Search("hello")
	fmt.Println("found: hello", "definition:", definition)
	err2 := dictionary.Add("hello", "Greeting")
	if err2 != nil {
		fmt.Println(err2)
	}
}