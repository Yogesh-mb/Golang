package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruit = []string{}
	fmt.Printf("Type of fruit is: %T\n", fruit)
	// fruit[0] = "apple"
	// fruit[1] = "orange"
	// fruit[3] = "watermelon"
	fruit = append(fruit, "mango", "banana", "watermelon", "apple")
	fmt.Println("Fruit list is: ", fruit)

	fruit = append(fruit[1:3]) //start from 1 and end before 3
	fmt.Println(fruit)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 271
	highScore[2] = 972
	highScore[3] = 372
	// highScore[4] = 548

	highScore = append(highScore, 478, 666, 845)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	// how to remove value from slice based on index
	var courses = []string{"nodejs", "golang", "java", "c++", "python"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
