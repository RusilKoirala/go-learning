package main

import "fmt"

type myEnum string

const (
	Recieved  myEnum = "Recieved"
	Confirmed        = "Confirmed"
	Prepared         = "Prepared"
	Delivered        = "Delivered"
)

func main() {
	fmt.Println("Enums")
	fmt.Println("Product status", Confirmed)
}
