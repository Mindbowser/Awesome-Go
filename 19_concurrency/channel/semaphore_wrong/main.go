package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	//we will block here until done becomes true
	go func() {
		<-done
		<-done
		close(c)
	}()
	//to unblock above we need to take off values of chan here
	//but we never got bcz we'r blocked already

	for n := range c {
		fmt.Println(n)
	}

}
