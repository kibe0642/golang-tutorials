package main

import "fmt"

func main() {
	fmt.Println("maps in golang")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all Languages:", languages)

	fmt.Println("JS is short for:", languages["JS"]) //Retrieving one input from a list
	fmt.Println("RB is short for:", languages["RB"])

	fmt.Println("PY is short for:", languages["PY"])

	//delete some values from a map

	delete(languages, "RB")
	fmt.Println("List of all Languages:", languages)

	//looping in maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}

//AUTHOR is Ras_aghul
