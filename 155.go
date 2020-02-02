package main

import (
	"fmt"
)

type MinStack struct {
	Vals []int
	Mins []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if len(this.Mins) == 0 || x <= this.Mins[len(this.Mins)-1] {
		this.Mins = append(this.Mins, x)
	} else if len(this.Mins) > 0 && x > this.Mins[len(this.Mins)-1] {
		this.Mins = append(this.Mins, this.Mins[len(this.Mins)-1])
	}
	this.Vals = append(this.Vals, x)
}

func (this *MinStack) Pop() {
	if len(this.Vals) < 1 {
		return
	}
	this.Vals = this.Vals[0 : len(this.Vals)-1]
	this.Mins = this.Mins[0 : len(this.Mins)-1]
}

func (this *MinStack) Top() int {
	if len(this.Vals) < 1 {
		return -1
	}
	return this.Vals[len(this.Vals)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Mins) < 1 {
		return -1
	}
	return this.Mins[len(this.Mins)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
