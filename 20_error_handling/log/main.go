package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nsf, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nsf)

}

func main() {
	_, err := os.Open("a.txt")
	if err != nil {
		log.Println("Error present there", err)
		fmt.Print(err)
	}

}
