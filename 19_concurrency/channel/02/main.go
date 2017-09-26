package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		fmt.Println("Blocked")
		close(ch)
	}()

	for n := range ch {
		time.Sleep(5 * time.Millisecond)
		fmt.Println("Printing")
		fmt.Println(n)
	}

}
