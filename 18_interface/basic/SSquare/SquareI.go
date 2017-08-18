package SSquare

import (
	"fmt"
)

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	fmt.Println(s.Side * s.Side)
	return s.Side * s.Side
}
