package structs

import "fmt"

// struct declaration
// capital letters = public visibility
type User struct {
	Name   string
	Age    Legal
	Gender string
	Weight float32
	Height float32
}

// (u User) - called 'receiver'
// Receiver by value
func (u User) PrintUserInfo() {
	fmt.Printf("Name: %s\nAge: %d\nGender: %s\nWeight: %f\nHeight: %f\n\n",
		u.Name, u.Age, u.Gender, u.Weight, u.Height)
}

func (u *User) PrintUserInfoWithPointer(name string) {
	u.Name = name
	fmt.Printf("Name: %s\nAge: %d\nGender: %s\nWeight: %f\nHeight: %f\n\n",
		u.Name, u.Age, u.Gender, u.Weight, u.Height)
}

// OR
// much easier
func (u User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u User) SetNameByValue(name string) {
	u.Name = name
}

// typecasting
type Legal int

func (l Legal) IsLegal() bool {
	return l >= 18
}

func CreateUser(name string, age int, gender string, weight float32, height float32) User {
	return User{
		Name:   name,
		Age:    Legal(age), // type conversion
		Gender: gender,
		Weight: weight,
		Height: height,
	}
}
