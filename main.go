package main

import "fmt"

func main() {
	messages := [3]int{0, 1, 2}

	defer handlePanic()

	// START HERE
	for i := range messages {
		messages[i+1]++
	}

	// DEFER (additional 50 ms)
	defer printSmth()

	fmt.Println("0")
	fmt.Println("1")
	fmt.Println("2")
	defer fmt.Println("test")
	fmt.Println("3")
	fmt.Println("4")
}

// RARE
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
		fmt.Println("handlePanic() successfully executed")
	}
}

func printSmth() {
	fmt.Println("smth...")
}
