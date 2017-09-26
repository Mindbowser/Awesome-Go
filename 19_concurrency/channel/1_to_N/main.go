package main

import (
	"fmt"
)

func main() {
	nums := 10
	c := make(chan int)
	ok := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("Receiving")
		close(c)
	}()

	for i := 0; i < nums; i++ { //writing to 10 functions at a times
		go func() {
			for n := range c {
				fmt.Println("Hi", n)
			}
			fmt.Println("Sending")
			ok <- true
		}()
	}

	for i := 0; i < nums; i++ {
		<-ok
	}
}
