package main

import "fmt"

func main() {
	//how to remove a value from slice based on index
	var Courses = []string{"Javascript", "Ruby", "Reactjs", "Swift", "Python"}

	fmt.Println(Courses)

	var index int = 2
	Courses = append(Courses[:index], Courses[index+1:]...)
	fmt.Println(Courses)
}
