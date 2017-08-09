package main

import (
	"fmt"
)

func main() {
	test("one")

}
func test(str string) {

	val := str
	switch val {
	case "zero", "Zero":
		fmt.Println("Zero:", val)
		fallthrough /*even if next condition fails
		it will enter into it*/
	case "one", "One":
		fmt.Println("One:", val)
		break
	default:
		fmt.Println("Wrong one !!")

	}
}
