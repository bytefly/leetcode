package main

import (
	"fmt"
)

type MyStack struct {
	Val []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{make([]int, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Val = append(this.Val, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	top := this.Val[len(this.Val)-1]
	this.Val = this.Val[:len(this.Val)-1]
	return top
}

/** Get the top element. */
func (this *MyStack) Top() int {
	top := this.Val[len(this.Val)-1]
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Val) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Empty())
}
