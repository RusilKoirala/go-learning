package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("[Slices]")

	// Basically an array without the size of array is a slice but with more functions

	var fruit = []string{}
	jeffFruit := append(fruit, "Banana", "Peach")

	fmt.Println("Fruit list : ", fruit)
	fmt.Println("Jeff fruit list : ", jeffFruit)

	// cool much simple :)

	// now lets do some slicing

	something := []int{1, 2, 3, 4, 5, 6}
	newList := append(something[:])
	fmt.Println(newList)
	// just play with it until you get it :)
	// it also works for string VERY cool

	// Another syntac to make array are
	scores := make([]int, 4)
	scores[0] = 3553
	scores[1] = 2343
	scores[2] = 6466
	scores[3] = 2345
	fmt.Println(scores)

	// so it also works even if i append above the bound of 4
	scores = append(scores, 5000)
	fmt.Println(scores)

	// i can also use sorting function on it like normal slices

	sort.Ints(scores)
	fmt.Println(scores)
}
