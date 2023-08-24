package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Server in Go")
	PerfomGetRequest()

}
func PerfomGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)

	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	/* if err!=nil{
		panic(err)
	} */
	fmt.Println(string(content))
}
