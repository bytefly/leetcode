package main

import (
	"fmt"
)

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{0, nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	i := 0
	p := this.Next
	for p != nil {
		if i == index {
			return p.Val
		}
		p, i = p.Next, i+1
	}

	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := &MyLinkedList{val, nil}
	head.Next = this.Next
	this.Next = head
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	p := this
	for p.Next != nil {
		p = p.Next
	}
	tail := &MyLinkedList{val, nil}
	p.Next = tail
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	i := 0
	p := this.Next
	for p != nil {
		if i == index-1 {
			q := &MyLinkedList{val, nil}
			q.Next = p.Next
			p.Next = q
		}
		p, i = p.Next, i+1
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	i := -1
	p := this
	for p != nil {
		if i == index-1 {
			q := p.Next
			if q == nil {
				p.Next = nil
			} else {
				p.Next = q.Next
			}
			break
		}
		p, i = p.Next, i+1
	}
}

func (this *MyLinkedList) GetValSlice() []int {
	vals := make([]int, 0)
	p := this.Next
	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}
	return vals
}

func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.DeleteAtIndex(1)
	fmt.Println(obj.GetValSlice())
}
