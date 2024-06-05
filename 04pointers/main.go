package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNum := 23
	var pointer = &myNum

	fmt.Println("Value of actual pointer is", pointer)
	fmt.Println("Value of actual pointer is", *pointer)
	*pointer = *pointer * 2
	fmt.Println("New value is", *pointer)
}
