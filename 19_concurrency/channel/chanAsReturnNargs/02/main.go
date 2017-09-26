package main

import "fmt"

func main() {
	c := send()
	cSum := receive(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func send() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func receive(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
