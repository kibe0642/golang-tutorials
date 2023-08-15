package main

import "fmt"

//main function ...entry point for the program in go
func main() {
	fmt.Println("Welcome to Functions") /*executes
	  automatically*/
	greetings() //calling function

	result := adder(3, 5)
	fmt.Println("sum of two numbers is:", result)

	proRes := proAdd(34, 73, 89, 637, 88)
	fmt.Println("Pro result:", proRes)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func proAdd(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

//called function
func greetings() {
	fmt.Println("Iamunee oo")

}
