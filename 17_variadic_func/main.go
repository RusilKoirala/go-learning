package main

import "fmt"

func main() {
	b := additems(3, 4, 5, 5)
	fmt.Println(b)
	fmt.Println("w", "wwow", 23)
}

func additems(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func showInterface(inter ...interface{}) interface{} {
	return inter
}
