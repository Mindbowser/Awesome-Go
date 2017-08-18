package main

import (
	"fmt"
	"scope/scopes/myPack"
)

var x string = "Welcome to Golang"

func main() {
	y := print(x)
	fmt.Println(y)
	z := myPack.MyName //MyName should start with Capital letter
	//beacause we are accessing it from another package
	fmt.Println(z)
	a := myPack.YourName
	fmt.Println(a)

	start := call()
	fmt.Println("Start:", start)

}

func print(s string) string {
	return s
}

func call() int {
	x := 10
	x++
	return x
}
