package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in go")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAddr(4, 8, 7, 6, 5)
	fmt.Println("ProResult is : ", proResult)
}

func greeter() {
	fmt.Println("Namaste from golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAddr(values ...int) int {
	total := 0
	for val := range values {
		total += values[val]
	}
	return total
}
