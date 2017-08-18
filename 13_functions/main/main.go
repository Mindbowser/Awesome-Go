package main

import (
	"fmt"
	"os"
)

func addition(number1 int, number2 int) int {
	return number1 + number2 //returning one value
}

func pass(value1 int, value2 int) (int, int) {
	return value1, value2
} //returning multiple values

//Variadic Function
func sum(numbers ...float64) (res float64) {
	for _, number := range numbers {
		res += number
	}
	return res
}

//Func as expression value  --> Closure
func decrement() func() int {
	x := 50
	return func() int {
		x--
		return x
	}
}

//Callbacks
func show(numbers []int, call func(int)) {
	for _, n := range numbers {
		call(n)
	}
}

//Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//Passing value
func main() {

	add := addition(1087, 3422)
	fmt.Println("Addition:", add)
	get, set := pass(123, 543)
	fmt.Println("Return Multiple Values:", get, set)
	data := []float64{3.3, 4.4, 5.5}
	fmt.Println(sum(1.1, 2.2, 3.3))
	fmt.Println(sum(data...))

	//Function as expression
	message := func() {
		fmt.Println("Good Morning User")
	}
	message() //Expression name as function call
	fmt.Printf("Type of Expression Fun:%T\n ", message)

	//Closure
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println("Sum is:", sum(435, 765))
	dec := decrement()
	fmt.Println("Decremented Value:", dec())

	show([]int{43, 76, 34}, func(n int) {
		fmt.Println(n)
	})

	//Recursion
	fmt.Println("Factorial of 7 using Recursion", factorial(7))

	//Defer
	f, _ := os.Open("log.txt")
	fmt.Println("Opening file.....")
	defer f.Close()
	fmt.Println("File closed")
}
