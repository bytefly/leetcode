package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	t := int(math.Sqrt(float64(area)))
	l, w := t, area/t

	for l*w != area {
		l++
		w = area / l
	}
	if l < w {
		l, w = w, l
	}
	return []int{l, w}
}

func main() {
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(5))
	fmt.Println(constructRectangle(6))
	fmt.Println(constructRectangle(7))
	fmt.Println(constructRectangle(8))
}
