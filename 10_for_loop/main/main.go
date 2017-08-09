package main

import (
	"fmt"
)

func main() {
	//Nested for loop
	for i := 1; i <= 5; i++ {

		for j := 1; j <= 10; j++ {
			fmt.Print("\t")
			fmt.Println(j * i)
		}
	}
	//for as a while loop
	k := 0
	for k < 10 {
		fmt.Println(k)
		if k >= 10 {
			break
		}
		k++
	}
	//for with no condition
	m := 0
	for {
		fmt.Println(m)
		if m >= 10 {
			break
		}
		m++
	}
	for i := 65; i <= 122; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
		//converted int type to string and []byte
	}
}
