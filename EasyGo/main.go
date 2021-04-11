package main

import (
	"fmt"
	"study-go/EasyGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}