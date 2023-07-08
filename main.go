package main

import "fmt"

func main() {
	// POINTERS
	a, b := 5, 10
	fmt.Printf("before call a: %d, b: %d\n\n", a, b)

	changeValue(a, b)
	fmt.Printf("changeValue() a: %d, b: %d\n", a, b)

	changeValueWithPointers(&a, &b)
	fmt.Printf("changeValueWithPointers() a: %d, b: %d\n\n", a, b)

	c := 15
	checkMemAddress(&c)

	// ARRAYS
	// 1-st option of an array declaration
	var arr0 [2]int
	arr0[0] = 1
	arr0[1] = 2
	// arr0[2] = 3 // error

	// 2-nd option of an array declaration
	arr1 := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println("arr0: ", arr0)
	fmt.Println("arr1: ", arr1)
	fmt.Println()
}

func changeValue(a int, b int) {
	temp := a
	a = b
	b = temp
}

func changeValueWithPointers(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func checkMemAddress(c *int) {
	fmt.Println("memory address of the variable: ", c)
	fmt.Println("memory address of the pointer", &c)
	fmt.Println("value of the variable", *c)
}
