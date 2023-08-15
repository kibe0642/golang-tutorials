package main

import "fmt"

func main() { //executes code LIFO...**executes deferred code in reversed order**
	defer fmt.Println("one")
	defer fmt.Println("Seven")
	defer fmt.Println("Ten")
	fmt.Println("Defer in go")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
