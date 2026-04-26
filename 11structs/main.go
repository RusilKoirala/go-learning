package main

import "fmt"

func main() {
	fmt.Println("[STRUCTS]")
	// no inheritance in golang so we have structs
	myUser := User{"Rusil", "ceo@rusil.me", 17, true}
	fmt.Println(myUser)
	fmt.Printf("%+v\n", myUser)
	fmt.Println(myUser.Name)

	myCar := Car{"Proshe", 600, true, 240000.99}
	fmt.Printf("Car Brand %v\nCar Price %v", myCar.Brand, myCar.Price)

}

type User struct {
	Name       string
	Email      string
	Age        int
	IsVerified bool
}

type Car struct {
	Brand string
	HP    int
	isICE bool
	Price float32
}
