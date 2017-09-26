package main

import (
	//	"fmt"
	//	"log"
	"os"
)

func main() {
	_, err := os.Open("fileOne.txt")
	if err != nil {
	//	fmt.Println("Error Happened", err)
	//	log.Fatalln("Error Happened", err)
		panic(err)
	}
}














