package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("[ MATH TIME ]")

	// randomnumber := rand.Int()

	// Small game
	for {
		// Initialization
		randomNum := rand.Intn(10) + 1
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter a number (1-10)")
		// fmt.Println(randomNum)

		// Reading input and checking error
		numberr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[Error] reading input:", err)
			return
		}

		// Converting string to int and checking for error
		number, err := strconv.Atoi(strings.TrimSpace(numberr))
		if err != nil {
			fmt.Println("[SYSTEM] Please enter a valid whole number")
			return
		}

		// Main logic you know :)
		if number == randomNum {
			fmt.Println("[SYSTEM] You win")
		} else {
			fmt.Printf("[SYSTEM] You lose, the number was %d\n", randomNum)
		}

		// Continue or break the loop
		fmt.Printf("Do you want to continue (Y/N)? ")
		cont, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[ERROR] reading input:", err)
			return
		}

		if strings.TrimSpace(cont) == "N" {
			fmt.Println("[SYSTEM] Thanks for playing")
			break
		}
	}
}
