package main

import (
	"fmt"
)

func main() {
	// map declaration
	characteristics := make(map[string]int)
	characteristics["age"] = 22
	characteristics["weight"] = 80
	characteristics["growth"] = 181

	// or

	//characteristics := map[string]int{
	//	"age":    22,
	//	"weight": 80,
	//	"growth": 181,
	//}

	fmt.Println(characteristics)

	fmt.Println()

	// get element by key
	age := characteristics["age"]
	fmt.Println(age)

	fmt.Println()

	// check if element exists
	weight, exists := characteristics["weight"]
	if exists {
		fmt.Println(weight)
	} else {
		panic("Element does not exist!")
	}

	fmt.Println()

	// iterate through map
	for key, value := range characteristics {
		fmt.Println(key, value)
	}

	fmt.Println()

	// add element to map
	characteristics["salary"] = 9999
	fmt.Println(characteristics)

	fmt.Println()

	// replace element in map
	characteristics["age"] = 54
	fmt.Println(characteristics)

	fmt.Println()

	// delete element from map
	delete(characteristics, "salary")
	fmt.Println(characteristics)
}
