package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	ok := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}

		ok <- true
	}()

	go func() {
		for i := 11; i <= 20; i++ {
			c <- i
		}
		ok <- true
	}()

	go func() {
		<-ok
		<-ok
		close(c)

	}()

	for n := range c {
		fmt.Println(n)
	}

}
