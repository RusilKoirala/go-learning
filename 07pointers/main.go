package main

import "fmt"

func main() {
	fmt.Println("[Pointers]")

	WOW := 9
	good(&WOW)
	fmt.Println(WOW)
	// var name string = "jeff"
	// var namePointer *string = &name
	// fmt.Printf("Type of %T\n", namePointer)
	// fmt.Printf("%v\n", *namePointer)
	// fmt.Println(namePointer)
	// wowPoint()
}

// func wowPoint() {
// 	var amazingtext string = "Step off"
// 	var pointerForAmazingtext *string = &amazingtext
// 	fmt.Println(amazingtext)
// 	fmt.Println(pointerForAmazingtext)
// 	fmt.Println(*pointerForAmazingtext)
// 	wow := &amazingtext
// 	fmt.Println(wow)
// 	fmt.Println(*wow)

// }
func good(WOW *int) {
	*WOW = 5
}
