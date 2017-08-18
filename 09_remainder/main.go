package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is Odd number")
		} else {
			fmt.Println(i, "is Even number")
		}
	}

}
