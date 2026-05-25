package main

import "fmt"

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// LIPO
type hypo[T any] struct {
	elements []T
}

func main() {
	myHypo := hypo[int]{
		elements: []int{1, 2, 3},
	}
	// nums := []int{1, 2, 3, 4, 5}
	// people := []string{"Bob", "Marley", "Chad"}
	// printSlice(nums)
	// printSlice(people)

	fmt.Println(myHypo)
}
