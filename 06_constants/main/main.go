package main

import (
	"fmt"
)

type Byte float64

//Single Value Initializatiob
const pi = 3.14

//MultiValue Initialization
const (
	userName         = "product"
	password         = "product@123" //String constant
	a                = iota          //2
	b                = iota          //3
	c                = 1 << iota     //16  (iota=4)
	d        float64 = iota * 5      //25  iota val inc by 1
)

func main() {

	func() {
		rad := pi * 7 * 7
		fmt.Println("Radius of Circle:", rad)
		fmt.Println("UserName:", userName)
		fmt.Println("Password", password)
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)

	}()
}
