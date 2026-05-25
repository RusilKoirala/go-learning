package main

import (
	"WOW/auth"
	"WOW/models"
	"fmt"
)

func main() {
	rusilCred := models.User{
		Username: "rusil",
		Password: "HackMe",
	}
	fmt.Println()
	auth.Login(rusilCred.Username, rusilCred.Password)
	// color.Red(rusilCred.Username)
}
