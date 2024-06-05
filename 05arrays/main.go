package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on arrays")

	var veg [5]string
	veg[0] = "Onion"
	veg[1] = "Brinjal"
	veg[3] = "potato"
	veg[4] = "raddish"

	fmt.Println("Fruint list is", veg)
	fmt.Println("Fruint list is", len(veg))

}
