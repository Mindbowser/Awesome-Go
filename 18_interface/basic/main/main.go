package main

import (
	"fmt"
	"interface/basic/IShape"
	"interface/basic/SCircle"
	"interface/basic/SSquare"
)

func info(s IShape.Shape) {
	fmt.Printf("Type: %T", s)
	fmt.Println()
	fmt.Println("Area:", s.Area())

}

func main() {
	s := &SSquare.Square{10}
	info(s)

	c := &SCircle.Circle{7}
	info(c)

}
