package main

import "fmt"

func main() {
	fmt.Println("[IF-ELSE]")

	LoginCount := 15

	if LoginCount < 20 {
		fmt.Println("GACHA")
	} else {
		fmt.Println("You are saved")
	}
	add(2, 4)
	OddorEven(5)

	/* if err != nil {

	} */

}
func add(a int, b int) {
	fmt.Println(a + b)
}
func OddorEven(num int) {
	if num%2 != 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
