package main

import "fmt"

func main() {
	fmt.Println("[ ARRAYS ]")

	// so pretty much we have to use to slice most of the time

	var myarr = [5]string{"Apple", "Mango", "Peach", "Guava", "Watermelon"}
	fmt.Println(myarr)
	fmt.Println(len(myarr)) // will give the the max array size not the real stored len
	fmt.Println(myarr[3])

	for i := 0; i < 5; i++ {
		fmt.Println(i, myarr[i])
	}
}
