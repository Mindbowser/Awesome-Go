package main

import (
	"fmt"
)

type wish struct {
	message string
	name    string
}

func Type(x interface{}) {

	switch x.(type) { //this one is assert
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case wish:
		fmt.Println("wish")
	default:
		fmt.Println("Unknown Type")
	}
}
func main() {
	Type(7)
	Type(wish{"Good Morning", "User"})
}
