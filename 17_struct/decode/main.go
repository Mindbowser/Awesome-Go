package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	Name    string
	Age     int
	address int //this field will not be exported
}

func main() {
	var p1 person
	reader := strings.NewReader(`{"Name":"James", "Age":24, "address":"New York"}`)
	json.NewDecoder(reader).Decode(&p1)

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.address)
}
