package something

import "fmt"

func SayHello() {
	fmt.Println("Hello GO!")
}

func sayBye() {
	fmt.Println("This is private(not exported)")
}
