package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Hello World")
	// fmt.Printf("Hello world")

	fmt.Printf("Server is running on localhost:3000\n")
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("GET /me/user\n")
	}
}
