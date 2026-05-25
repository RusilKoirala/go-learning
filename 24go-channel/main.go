package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	fmt.Println("Processing number", <-numChan)
}

func main() {
	// messageChan := make(chan string)

	// messageChan <- "ping" //blocking

	// text := <-messageChan
	// fmt.Println(text) // Deadlock

	numChan := make(chan int)

	go processNum(numChan)

	numChan <- 5

	time.Sleep(time.Second * 2)

}
