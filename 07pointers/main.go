package main

import "fmt"

func main() {
	fmt.Println("[Pointers]")
	var name string = "jeff"
	var namePointer *string = &name
	fmt.Printf("Type of %T\n", namePointer)
	fmt.Printf("%v\n", *namePointer)
	fmt.Println(namePointer)
}
