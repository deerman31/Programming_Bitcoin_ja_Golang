package main

import (
	"fmt"
	"log"
	"math"
)

type Point struct {
	x, y, a, b int
}

func NewPoint(x, y, a, b int) (Point, error) {
	fx := float64(x)
	fy := float64(y)
	fa := float64(a)
	fb := float64(b)

	if math.Pow(fy, 2) != (math.Pow(fx, 3) + fa*fx + fb) {
		return Point{0, 0, 0, 0}, fmt.Errorf("(%d, %d) is not on the curve", x, y)
	}
	return Point{x, y, a, b}, nil
}

func main() {
	p1, err := NewPoint(-1, -1, 5, 7)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	p2, err := NewPoint(18, 77, 5, 7)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println(p1 != p2)
}
