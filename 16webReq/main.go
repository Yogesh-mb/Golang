package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://courses.learncodeonline.in/"

func main() {
	fmt.Println("LCO Web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T", response)
	defer response.Body.Close() //callers responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data", string(databytes))
}
