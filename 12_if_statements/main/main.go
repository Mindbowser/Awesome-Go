package main

import (
	"fmt"
)

func main() {
	no := 56
	if true {
		fmt.Println("This will work")
	}
	if false {
		fmt.Println("This will not work")
	}
	if !true {
		fmt.Println("This will not work")
	}
	if !false {
		fmt.Println("This will work")
	}

	n := 45
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	//If with short statement
	if computer := "Dell"; no > 50 {
		fmt.Println(computer)
	}

}
