package main

import (
	"fmt"
)

// WORKS BEFORE main() FUNCTION
func init() {
	fmt.Print("\nCheck if init() function works\n")
}

func main() {
	// 0
	printSmth("Test...\n")

	// 1
	fmt.Println(calculateRectArea(5, 5), "\n")

	// 2
	test0 := explainSprintf("Tony", 22)
	fmt.Println(test0 + "\n")

	// 3
	name, age := returnMultipleParameters("Tony", 22)
	fmt.Println(name, age, "\n")

	// 4
	fmt.Println(explainParams("Yeah, boy...", " It's me.\n"))

	// 5 - Closures
	closure := increment()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
}

// function declaration
func printSmth(message string) {
	fmt.Println(message)
}

// function, returning something
func calculateRectArea(a int, b int) int {
	return a * b
}

// function Sprintf() explanation
// this function generates a string and it can be placed to variable
func explainSprintf(name string, age int) string {
	return fmt.Sprintf("Hello, %s! Congratulations with %d years old!", name, age)
}

// function, returning multiple parameters
func returnMultipleParameters(name string, age int) (string, bool) {
	var result string
	if age < 18 {
		result = fmt.Sprintf("You are not legal yet.")
		return result, false
	}
	//else if age >= 54 {
	//	result = fmt.Sprintf("Easy there...")
	//	return result, errors.New("Sorry, you are too old!")
	//}
	result = fmt.Sprintf("Welcome aboard!")
	return result, true
}

// Params
func explainParams(strings ...string) string {
	var res string
	//for i := 0; i < len(strings); i++ {
	//	res += strings[i] + " "
	//}
	for i := range strings {
		res += strings[i] + " "
	}
	return res
}

// Closure
func increment() func() int {
	var counter int
	return func() int {
		counter += 5
		return counter
	}
}
