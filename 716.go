package main

import "fmt"

type MaxStack struct {
	Vals []int
	Maxs []int
}

func Constructor() MaxStack {
	var obj MaxStack
	return obj
}

func (this *MaxStack) Push(x int) {
	this.Vals = append(this.Vals, x)
	if len(this.Maxs) == 0 {
		this.Maxs = append(this.Maxs, x)
	} else {
		curMax := this.Maxs[len(this.Maxs)-1]
		if curMax > x {
			this.Maxs = append(this.Maxs, curMax)
		} else {
			this.Maxs = append(this.Maxs, x)
		}
	}
}

func (this *MaxStack) Pop() int {
	if len(this.Vals) > 0 {
		t := this.Vals[len(this.Vals)-1]
		this.Vals = this.Vals[:len(this.Vals)-1]
		this.Maxs = this.Maxs[:len(this.Maxs)-1]
		return t
	}
	return -1
}

func (this *MaxStack) Top() int {
	if len(this.Vals) > 0 {
		return this.Vals[len(this.Vals)-1]
	}
	return -1
}

func (this *MaxStack) PeekMax() int {
	if len(this.Maxs) > 0 {
		return this.Maxs[len(this.Maxs)-1]
	}
	return -1
}

func (this *MaxStack) PopMax() int {
	var stack []int
	if len(this.Maxs) == 0 {
		return -1
	}

	ans := this.Maxs[len(this.Maxs)-1]
	for {
		s := this.Vals[len(this.Vals)-1]
		this.Vals = this.Vals[:len(this.Vals)-1]
		this.Maxs = this.Maxs[:len(this.Maxs)-1]
		if s == ans {
			break
		}
		stack = append(stack, s)
	}

	for i := len(stack) - 1; i >= 0; i-- {
		x := stack[i]
		this.Vals = append(this.Vals, x)
		if len(this.Maxs) == 0 {
			this.Maxs = append(this.Maxs, x)
		} else {
			curMax := this.Maxs[len(this.Maxs)-1]
			if curMax > x {
				this.Maxs = append(this.Maxs, curMax)
			} else {
				this.Maxs = append(this.Maxs, x)
			}
		}
	}

	return ans
}

func main() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(4)
	obj.Push(3)
	obj.Push(2)
	obj.Push(1)
	fmt.Println(obj.PopMax())
	fmt.Println(obj.Top())
	fmt.Println(obj.PeekMax())
}
