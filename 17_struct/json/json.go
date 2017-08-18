package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname   string
	Lastname    string
	Age         int `json:"Umar"`
	notexported int //will not be exported
}

func main() {
	p1 := person{"Rob", "Pike", 24, 007}

	bs1, _ := json.Marshal(p1)
	fmt.Println(string(bs1))

	bs := []byte(bs1)
	json.Unmarshal(bs, p1)
	fmt.Println(bs)

}
