package main

import (
	"fmt"
)

type MyQueue struct {
	Input  []int
	Output []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{make([]int, 0), make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Input = append(this.Input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.Empty() {
		top := this.Output[len(this.Output)-1]
		this.Output = this.Output[:len(this.Output)-1]
		return top
	}

	return -1
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.Empty() {
		top := this.Output[len(this.Output)-1]
		return top
	}
	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Output) == 0 && len(this.Input) > 0 {
		for i := len(this.Input) - 1; i >= 0; i-- {
			this.Output = append(this.Output, this.Input[i])
		}
		this.Input = this.Input[:0]
	}
	return len(this.Output) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}
