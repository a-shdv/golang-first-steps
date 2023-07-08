package main

import (
	"fmt"
	"go-first-steps/structs"
	"reflect"
)

func main() {
	// struct initialization
	user := structs.User{
		Name:   "Tony",
		Age:    22,
		Gender: "male",
		Weight: 80.24,
		Height: 181.44,
	}

	// print struct
	user.PrintUserInfo()
	user.PrintUserInfoWithPointer("unknown")
	user.PrintUserInfo()

	fmt.Println()

	fmt.Printf("%+v\n\n", user)

	// Receiver by value and by reference
	fmt.Printf("Name: %s\n", user.GetName())   // value
	user.SetName("Tony")                       // reference
	fmt.Printf("Name: %s\n\n", user.GetName()) // value

	// Receiver by value and by reference
	fmt.Printf("Name: %s\n", user.GetName()) // value
	user.SetNameByValue("Unknown")           // value (makes a copy), hence, doesn't change the actual value
	fmt.Printf("Name: %s\n", user.GetName()) // value

	// type conversion
	anotherUser := structs.User{
		Name:   "Alien",
		Age:    5,
		Gender: "unknown",
		Weight: 80.24,
		Height: 181.44,
	}
	isLegal := anotherUser.Age.IsLegal()
	fmt.Println(isLegal)
	fmt.Println(anotherUser.Age)
	fmt.Println(reflect.TypeOf(anotherUser.Age))
}
