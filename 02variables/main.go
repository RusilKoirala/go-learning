package main

import "fmt"

// jefftruename := "jeffbezozzz" [ PLS DONT UNCOMMENT IT *-*]
var jefftruename string = "jeffbezozozo"

const jeffage int = -1 // YES HE IS IMMORTALLL

func main() {
	fmt.Println("[Variables]")

	// Designing animal

	const dna string = "XB2E4" // dna := "XB2E4"
	const name string = "jeff" // name := "jeff"
	const isloyal bool = true  // isloyal := true
	var age int = 5            // age := 5

	// fmt.Println(dna, name, age)
	fmt.Printf("DNA : %s \nName : %s \nAge : %v\nLoyality : %v \n", dna, name, age, isloyal)
	fmt.Printf("dna datatype is %T\n", dna)

	// Types of int
	var smallInt uint8 = 255 // try to increment by one :)
	fmt.Println(smallInt)

	// floats
	const π float32 = 3.14159

	// Defualt values
	var somethingint int
	var somethingfloat float32
	var somethingbool bool
	var somethingstring string
	fmt.Println(somethingint, somethingbool, somethingfloat, somethingstring, jeffage)

}
