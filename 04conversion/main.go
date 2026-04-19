package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("[CONVERSION]")
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Magesty tip me \n")

	input, _ := reader.ReadString('\n')
	numinp, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// ERROR HANDLING FOR JEFFFF
	if err != nil {
		fmt.Println(err)
	}

	// GAME NAME: JEFF THE BEGGAR
	if numinp > 100 {
		fmt.Printf("[SYSTEM]: Transaction alert -%v from account\n", numinp)
		fmt.Printf("[JEFF] : THANK YOU\n")
	} else {
		fmt.Printf("[JEFF] : STEPP OFF BITCH\n")
	}

}
