package main

import (
	"fmt"
	"math"
)

func ccc(a float64) float64 {
	r := a
	var x = r * r
	return x
}

func main() {
	f := math.Pi
	g := f * ccc(2)
	fmt.Println(g)
}
