package main

import (
	"fmt"
)

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		for n := range c {
			out <- fact(n)
		}

		close(out)
	}()

	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total = total * i

	}

	return total
}

func main() {

	c := gen()
	d := factorial(c)
	for s := range d {
		fmt.Println(s)
	}

}
