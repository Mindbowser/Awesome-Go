package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person /*This type of []Person has
implemented sort.Interface interface*/

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age > b[j].Age
}

func main() {

	persons := []Person{
		{"George", 34},
		{"Robin", 43},
		{"Stuart", 23}}

	peoples := []string{"Keyboard", "mouse", "Computer"}
	rev := []string{"Book", "Literature", "Novel"}
	ints := []int{54, 76, 33, 12, 54, 76}
	fmt.Println("Before Sort:", persons)
	fmt.Println("Before Sort:", peoples)
	sort.Strings(peoples) /* Sorting string slice */
	sort.Sort(sort.Reverse(sort.StringSlice(rev)))
	fmt.Println("Reversed String Slice", rev)
	sort.Sort(sort.IntSlice(ints))
	fmt.Println("Sorted Integer value", ints)
	fmt.Println("After Sort:", peoples)
	sort.Sort(ByAge(persons))
	fmt.Println("After Sort:", persons)

}
