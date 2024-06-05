package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.youtube.com:3000/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	// parsing the url
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params is %T\n", qparams)

	fmt.Println(qparams["index"])

	for _, val := range qparams {
		fmt.Println("Param is ", val)
	}

	//building a url when you have all parts
	partsofUrl := &url.URL{ //& is important
		Scheme:   "https",
		Host:     "10.159.111.111:3001",
		Path:     "haha",
		RawPath:  "user=yogesh",
		RawQuery: "a=happy&b=sad",
	}

	anotherUrl := partsofUrl.String()

	fmt.Println(anotherUrl)

}
