package main

import "fmt"

func main() {
	go show()
	go display()
}

func show() {
	for i := 0; i < 45; i++ {
		fmt.Println("show:", i)
	}
}

func display() {
	for i := 0; i < 45; i++ {
		fmt.Println("display:", i)
	}
}
