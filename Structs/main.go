package main

import "fmt"

func main() {

	fmt.Println("welcome to structs")

	Aghul := Ras{"Aghul", "Aghulras@gmail.com", true, 23}

	fmt.Println(Aghul)
	fmt.Printf("My details are: %+v\n", Aghul)
}

//defining struct in golang
//notice structs Ras has first letter capitalized
type Ras struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
