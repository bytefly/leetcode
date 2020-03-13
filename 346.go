package main

import "fmt"

type MovingAverage struct {
	Size int
	Vals []int
	Sum  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size, nil, 0}
}

func (this *MovingAverage) Next(val int) float64 {
	this.Sum += val
	this.Vals = append(this.Vals, val)
	if len(this.Vals) <= this.Size {
		return float64(this.Sum) / float64(len(this.Vals))
	}
	this.Sum -= this.Vals[0]
	this.Vals = this.Vals[1:]
	return float64(this.Sum) / float64(this.Size)
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj.Next(1))
	fmt.Println(obj.Next(10))
	fmt.Println(obj.Next(3))
	fmt.Println(obj.Next(5))
}
