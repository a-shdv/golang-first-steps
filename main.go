package main

import (
	"fmt"
	"go-first-steps/shapes"
)

func main() {
	circle := shapes.Circle{Radius: 12.25}
	square := shapes.Square{Length: 12.25}
	printArea(circle)
	printArea(square)
}

// Принимает любой объект
func printArea(s shapes.Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", s.Area())
}
