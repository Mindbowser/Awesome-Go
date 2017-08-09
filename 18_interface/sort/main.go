package main

import (
	"fmt"
	"sort"
	//"strings"
)

func main() {

	s := []string{"John", "Al", "Zenny"}
	sort.StringSlice.Sort(s)
	fmt.Println(s)

	//Reverse of slice of strings

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	i := sort.Reverse(sort.StringSlice(s))
	fmt.Println(i)
}
