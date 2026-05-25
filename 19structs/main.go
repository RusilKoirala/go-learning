package main

import (
	"fmt"
	"time"
)

type order struct {
	name         string
	price        int
	isOutofStock bool
	createdAt    time.Time
}

func main() {
	wow := order{
		name:         "Hit",
		price:        50,
		isOutofStock: true,
		createdAt:    time.Now().UTC(),
	}
	fmt.Println(wow)
}
