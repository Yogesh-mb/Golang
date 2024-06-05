package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	performPerformFormRequest()

}

//get request
func PerformGetRequest() {
	const myURL = "http://localhost:8000/get"
	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code of resp: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(string(content))
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())
}

//post request
func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename": "lets go with golang",
		"price" : 0,
		"platform": "youtube.com
	} 
	`)

	resp, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(contents))
}

func performPerformFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "yogesh")
	data.Add("secondname", "badagi")
	data.Add("email", "abc@hsg")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(contents))
}
