package main

import "fmt"

func main() {
	fmt.Println("[Maps]")

	computerHardware := make(map[string]string)
	computerHardware["CPU"] = "For processing, scheduling tasks"
	computerHardware["RAM"] = "For storing short term data"
	computerHardware["GPU"] = "For performing small mathematical operation over large number of small computers"

	fmt.Println(computerHardware)

}
