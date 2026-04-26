package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("[SWITCH CASE]")

	rand.Seed(time.Now().Unix())
	DiceNumber := rand.Intn(6) + 1
	switch DiceNumber {
	case 1:
		fmt.Println("Dice 1 : You can move")
	case 2:
		fmt.Println("Dice 2 : You can move 2")
	case 3:
		fmt.Println("Dice 3 : You can move 3")
	case 4:
		fmt.Println("Dice 4 : You can move 4")
	case 5:
		fmt.Println("Dice 5 : You can move 5")
	case 6:
		fmt.Println("Dice 6 : You can move 6")
	default:
		fmt.Println("ERROR")
	}

}
