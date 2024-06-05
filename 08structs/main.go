package main

import "fmt"

func main() {
	fmt.Println("Details on structs")

	yogesh := User{"Yogesh", "abc@xyz", true, 58}
	fmt.Println(yogesh)
	fmt.Printf("Yogesh details are %+v\n", yogesh)
	fmt.Printf("Name is %v and emal is %v\n", yogesh.Name, yogesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
