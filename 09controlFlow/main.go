package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional statements in golang")

	count := 54
	var result string

	if count < 10 {
		result = "nice"
	} else if count > 10 {
		result = "ok"
	} else {
		result = "not nice"
	}

	fmt.Println(result)

	if 5%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 7; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}
}
