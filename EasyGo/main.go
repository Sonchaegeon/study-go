package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upperName := lenAndUpper("sonchaegeon")
	fmt.Println(multiply(2, 4))
	fmt.Println(totalLength, upperName)
	repeatMe("son", "go", "park", "lee", "kim")
}
