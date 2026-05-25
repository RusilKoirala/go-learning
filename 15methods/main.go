package main

import "fmt"

func main() {
	fmt.Println("[METHODS]")

	rusil := User{"Rusil", 17, "rusil@me", true, true}
	rusil.getStatus()
	fmt.Println(rusil.Email)
	rusil.resetEmail()
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
	isCool bool
}

func (u User) getStatus() {
	fmt.Println("User is active : ", u.Status)
}

func (e User) resetEmail() {
	e.Email = "rusil@go"
	fmt.Println("New email ", e.Email)
}
