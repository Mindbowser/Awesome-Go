package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{"George", 23}

	json.NewEncoder(os.Stdout).Encode(p1) //Encoding through streams

	fmt.Println(p1.Name, "--", p1.Age)

}
