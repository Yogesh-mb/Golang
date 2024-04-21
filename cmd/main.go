package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	// if len(os.Args) > 1 {
	// fmt.Println(hello.Say(os.Args[1]))
	// } else {
	// 	fmt.Println(hello.Say("Hello, World!"))
	// }

	fmt.Println(hello.Hey(os.Args[1:]))
}
