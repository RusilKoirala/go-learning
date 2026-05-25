package main

import (
	"fmt"
)

func main() {
	fmt.Println("[LOOPS]")

	// Old way :)
	for i := 0; i < 4; i++ {
		println(i, "HI")
	}
	// New way
	for i := range 5 {
		fmt.Println(i, "HI")
	}

	MySLICE := []string{"APPLE", "MANGO", "STRAWBERRY", "GRAPES"}

	for i, day := range MySLICE {
		fmt.Println(i, MySLICE[i], day)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

}
