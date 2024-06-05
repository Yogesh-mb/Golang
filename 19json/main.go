package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to class on JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS bootcamp", 199, "lco.org", "hagt@13", []string{"web-dev", "js"}},
		{"MERN bootcamp", 599, "lco.org", "hagt@56", []string{"full-stack", "js"}},
		{"ANgular bootcamp", 299, "lco.org", "hakh@13", []string{"facebook", "js"}},
	}

	//package the above data into json

	finalJson, err := json.MarshalIndent(lcoCourses, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS bootcamp",
		"Price": 199,
		"website": "lco.org",
		"tags": ["web-dev","js"]
	   }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json is not valid")
	}

	//cases where you want to add data in key value format
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", key, val, val)
	}

}
