package main

import (
	"errors"
	"fmt"
)

func main() {
	// SLICE DECLARATION (no arr size presented)
	// messages := []string{"1", "2", "3"}
	// OR
	// 1-st arg: slice type
	// 2-nd arg: slice len
	// 3-rd arg: slice capacity
	messages := make([]string, 2, 5)
	// messages := make([]string, 5) fills in the slice with 5 empty elements
	messages[0] = "test0"
	messages[1] = "test1"
	// messages[2] = "test1" // error

	// adds elements to the end of the slice
	// and automatically increases 2x times slice capacity
	messages = append(messages, "test2", "test3")

	fmt.Println(messages)
	fmt.Println("length: ", len(messages))
	fmt.Println("capacity:", cap(messages))

	fmt.Println()

	// ERROR HANDLING
	errMessage, err := isEmpty(messages)
	if err != nil {
		fmt.Println(errMessage, err)
	} else {
		fmt.Println(errMessage)
	}
}

func isEmpty(str []string) (string, error) {
	if len(str) == 0 {
		return "Error!", errors.New("Array is empty!")
	}
	return "Success!", nil
}
