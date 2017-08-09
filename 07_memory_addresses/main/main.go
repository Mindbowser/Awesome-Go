package main

import (
	"fmt"
)

const pi = 3.14

var message string = "Product"

func main() {
	responseCode := 101
	fmt.Printf("\n message:%v,Address:%v", message, &message)
	fmt.Printf("\n responseCode:%v,Address:%v", responseCode, &responseCode)
	area := pi * 7 * 7
	fmt.Printf("\n Area:%v,Address:%v", area, &area)

}
