package main

import (
	"fmt"
)

func main() {

	a := 50
	fmt.Println(a)
	fmt.Println("a:", &a) //Address of a
	/*Here x will point to the address where a
	is stored */
	var x = &a //Referencing
	/*Here x will get the value by using '*' */
	fmt.Println("x:", *x) //DeReferencing

	/*s := 5
	call(s)
	fmt.Println(s)*/

	s := 5
	fmt.Println("Address of S:", &s)
	call(&s)
	fmt.Println(s)

}

func call(n *int) {
	fmt.Println("Address:", &n)
	*n++

}

/*func call(n int) {

	n++
}*/
