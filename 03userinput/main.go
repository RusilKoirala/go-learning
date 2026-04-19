package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("[User Input]")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for JEFF (MY BOII)")

	// comma ok syntax || err ok  WOW

	input, _ := reader.ReadString('\n')

	/* So baiscally, the comma ok syntax works like try-catch of js. If there is
	no error then the value gets stored in (input) variable. But if there is an
	error the error is sent to the second variable :)

	But we can put _ as nothing (i guess) the go compiler dont seems to worry
	about it being used or not*/

	fmt.Printf("Type of %T\n", input)
	fmt.Println("DID YOU RATE JEFF", input)

}
