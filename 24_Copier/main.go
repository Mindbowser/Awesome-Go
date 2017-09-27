package main
import (
	"fmt"
	"github.com/jinzhu/copier"  )
type User struct {
	Name string
	Age  int32
}
type Employee struct {
	Name      string
	Age       int32
	DoubleAge int32
	EmployeId int64
}
func main() {
	var (
		user      = User{Name: "John", Age: 18}
		users     = []User{{Name: "Martin", Age: 18}, 
						   {Name: "Ricky", Age: 30}}
		employee  = Employee{}
		employees = []Employee{}
	)
	//Copy field to field
	copier.Copy(&employee, &user)
	fmt.Println(employee)
	// Copy struct to slice
	copier.Copy(&employees, &user)
	fmt.Println(employees)
	// Copy slice to slice
	employees = []Employee{}
	copier.Copy(&employees, &users)
	fmt.Println(employees)	
}