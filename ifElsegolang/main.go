package main

import "fmt"

func main() {
	fmt.Println("if else in go")
	LoginCounter := 10
	var result string
	if LoginCounter < 5 {
		result = "Regular User"
	} else if LoginCounter > 5 {
		result = "Watch out"
	} else {
		result = "exactly count 5"
	}
	fmt.Println(result)
}
