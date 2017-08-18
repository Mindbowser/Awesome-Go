package main

import "fmt"

type width int //user defined type
type Mobile struct {
	brand string
	model string
	price int
}

func (mob Mobile) display() string { //func associated with Mobile
	mob.brand = "Xiomi"
	return mob.brand
}
func (mob *Mobile) show() string { //func associated with Mobile
	mob.brand = "Xiomi"
	return mob.brand
}
func main() {
	var height width
	fmt.Println(height)
	m := Mobile{}
	fmt.Println(m) //Default values inside struct{" " 0}
	var mob Mobile //Instance creation using var
	fmt.Println(mob)
	mobs := new(Mobile)
	fmt.Println(mobs)
	phone := Mobile{"Samsung", "Galaxy", 24500}
	fmt.Println("Before Change:", phone)
	fmt.Println("Function Call", phone.display()) // Xiomi
	fmt.Println("After Change:", phone)           /*still old values are coming
	{"Samsung","Galaxy",24500}*/
	fmt.Println("Function Call:", phone.show()) //calling show()
	fmt.Println("After Change:", phone)         //Here changed values will reflect
	phones := &Mobile{"Lenovo", "A6", 14300}
	fmt.Println(phones.model)
}
