package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch case")
	//generate random numbers

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Number on the dice:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice num,ber is 1 and you can open")
	case 2:
		fmt.Println("you can move spot 2")
	case 3:
		fmt.Println("You can move to spot 3")
		fallthrough
	case 4:
		fmt.Println("You can move to spot 4")
	case 5:
		fmt.Println("You can move to spot 5")
		fallthrough
	case 6:
		fmt.Println("You can move to spot 6 and roll the dice again")
	default:
		fmt.Println("What was that")
	}

}
