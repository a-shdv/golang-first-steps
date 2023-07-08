package main

import (
	"fmt"
	"reflect"
)

func main() {
	// TYPE DECLARATIONS
	{
		// 1-st option to declare a variable
		var message0 string
		message0 = "Test..."

		// 2-nd option to declare a variable
		message1 := "Test..."

		fmt.Println("message0: " + message0)
		fmt.Println("message1: " + message1)
		fmt.Println()

		// null value
		var message2 int
		fmt.Print("message2: ")
		fmt.Println(message2) // returns 0
		fmt.Println()

		// if you want to know the type of variable, then you can use reflect.TypeOf(var)
		fmt.Print("type of message1: ")
		fmt.Println(reflect.TypeOf(message1))
		fmt.Println()

		// multiple initialization
		a, b, c := 1, 2, 3
		fmt.Printf("a: %d, b: %d, c: %d", a, b, c)
		fmt.Println()

		// swap variables
		a, b = b, a
		fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
		fmt.Println()
	}

	// BASIC TYPES
	{
		// byte
		var bvar byte
		bvar = 255 // 0..255

		// int
		var ivar int
		ivar = 1225

		// float
		var fvar float32
		fvar = 12.25

		// boolean
		var boolvar bool
		boolvar = true

		// character
		var rvar rune // the same as 'char' type
		rvar = 'a'

		// string
		var svar string
		svar = "abcdef..."

		fmt.Println(bvar)
		fmt.Println(ivar)
		fmt.Println(fvar)
		fmt.Println(boolvar)
		fmt.Printf("char: %c, code: ", rvar) // the same as in C
		fmt.Println(rvar)
		fmt.Println(svar)
	}
}
