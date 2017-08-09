package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (per *Person) show() {
	fmt.Println(per.FirstName, "", per.LastName, "", per.Age)
}

func (cust *Customer) show() { //method overriding based on receiver
	fmt.Println(cust.ContactNo)
}

type Customer struct {
	Person
	ContactNo int
}

func main() {
	cust := Customer{
		Person: Person{
			"John",
			"Matthew",
			25,
		},
		ContactNo: 9965,
	}
	cust.Person.show()
	cust.show()
}
