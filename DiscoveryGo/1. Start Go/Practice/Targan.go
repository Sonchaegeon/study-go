package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 10; i <= n * 10; i += 10 {
		fmt.Printf("타잔이 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", i, i + 10)
	}
}
