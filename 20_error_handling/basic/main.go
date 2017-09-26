package main

import (
	"fmt"
	"log"
)

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Negative Number==>%v", n)
	}
	return n * n, nil
}
func main() {
	a, err := Sqrt(-21)
	if err != nil {
		fmt.Println("Error occured", err)
		log.Fatal(err)
	}
	fmt.Println(a)
}
