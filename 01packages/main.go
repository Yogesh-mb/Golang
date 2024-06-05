package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pudding")

	//comma ok syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

}
