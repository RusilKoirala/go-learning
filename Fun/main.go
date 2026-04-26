package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("[COMPUTER]")
	num := runtime.NumCPU()
	versionn := runtime.Version()
	fmt.Println("CPU cores", num)
	fmt.Println(versionn)
	fmt.Println("Hello")
	x := 0
	for {
		fmt.Println(x)
		if (x > 5 ) {
			break;
		}
		x++;
	}

	add (5,3)
}

func add (int a , int b) {
	Println("The sum is", a+b)
}
