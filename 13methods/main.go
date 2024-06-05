package main

import "fmt"

func main() {
	fmt.Println("Details on methods")

	yogesh := User{"Yogesh", "abc@xyz", true, 58}
	fmt.Println(yogesh)
	fmt.Printf("Yogesh details are %+v\n", yogesh)
	fmt.Printf("Name is %v and emal is %v\n", yogesh.Name, yogesh.Email)
	yogesh.GetStatus()
	yogesh.NewMail()
	fmt.Printf("Name is %v and emal is %v\n", yogesh.Name, yogesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@mail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
