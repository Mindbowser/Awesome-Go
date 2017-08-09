package main

import (
	"fmt"
)

func display(v interface{}) {
	fmt.Printf("value is:%v", v.(int))
}

func main() {
	v := 10
	display(v)
}
