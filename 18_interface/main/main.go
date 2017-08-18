package main

import (
	"fmt"
)

type vehicles interface{}

type vehicle struct {
	speed string
	color string
}

type car struct {
	vehicle
	wheels int
	doors  int
}

type plane struct {
	vehicle
	jet bool
}

type boat struct {
	vehicle
	capacity int
}

func specifications(veh interface{}) { /* accepting anything*/
	fmt.Println(veh)
}

func main() {
	audiR8 := car{vehicle{"240 KMPH", "White"}, 4, 4} /* specify values*/
	benz := car{vehicle{"250 KMPH", "Black"}, 4, 4}
	boeing := plane{vehicle{"500 KMPH", "White"}, true}
	sanger := boat{vehicle{"350 KMPH", "Saffron"}, 10000}
	travelling := []vehicles{audiR8, benz, boeing, sanger}

	for key, value := range travelling {
		fmt.Println(key, "-", value)
	}

	specifications(audiR8)
	specifications(benz)
	specifications(boeing)
	specifications(sanger)

	var c rune = 'a'
	var y int32 = 'b'
	fmt.Println(string(c))
	fmt.Println(string(y))

}
