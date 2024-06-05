package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on maps")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"
	fmt.Println("lIST OF LANGUAGES", languages)

	delete(languages, "py")
	fmt.Println("lIST OF LANGUAGES", languages)

	// loops
	for key, val := range languages {
		fmt.Printf("For key %v, value is %v\n", key, val)
	}

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
