package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func main() {
	c := make(chan int)
	w.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		w.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		w.Done()
	}()

	go func() {
		w.Wait()
		close(c)
		fmt.Println("Channel closed")
	}()

	for n := range c {
		fmt.Println(n)
	}

}
