package main

import (
	"fmt"
	"runtime"
	"sync"
)

var w sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(1)

	w.Add(2)
	fmt.Println("Starting goroutines")

	go func() {
		//defer w.Done()
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Println(char)
		}
		w.Done()
	}()

	go func() {
		//defer w.Done()
		for i := 0; i < 27; i++ {
			fmt.Println(i)
		}
		w.Done()
	}()
	w.Wait()

	fmt.Println("Waiting to finish program")
	fmt.Println("Program finished")
}
