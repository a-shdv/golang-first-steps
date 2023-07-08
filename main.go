package main

import "fmt"

func main() {
	// 2D array
	matrix := make([][]int, 10)

	// C-like loops
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j] = make([]int, 10) // allocate memory for inner slice
		}
	}

	// Golang style
	for i := range matrix {
		for j := range matrix {
			matrix[i][j] = j
		}
	}

	// While
	//for true {
	//	fmt.Println(matrix)
	//}

	fmt.Println(matrix)

	// Key-value
	for key, value := range matrix {
		fmt.Println(key, value)
	}
}
