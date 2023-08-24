package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Password string   `json:"-"`
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Create Json in go")
	EncodeJson()
}
func EncodeJson() {
	rascourses := []course{
		{"React-bootcamp", 299, "ras123", "windows", []string{"Backend-dev", "go"}},
		{"Angular-bootcamp", 199, "ras8823", "windows", []string{"Backend-dev", "go"}},
		{"React-bootcamp", 239, "rase9123", "Unix", nil},
	}

	//package json
	finalJson, err := json.MarshalIndent(rascourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
