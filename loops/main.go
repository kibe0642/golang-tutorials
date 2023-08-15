package main

import "fmt"

func main() {

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	fmt.Println(days)

	myValue := 1
	for myValue < 10 {
		/* if myValue == 5 {
			break
		} */
		if myValue == 7 {
			myValue++
			continue
		}
		if myValue == 9 {
			goto ras
		}
		fmt.Println(myValue)
		myValue++
	}
ras:
	fmt.Println("I'm happy my code is working")

}
