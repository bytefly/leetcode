package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := 0, 0
	p := l1
	for p != nil {
		num1 *= 10
		num1 += p.Val
		p = p.Next
	}

	p = l2
	for p != nil {
		num2 *= 10
		num2 += p.Val
		p = p.Next
	}

	num := num1 + num2
	p = &ListNode{0, nil}
	for num > 0 {
		q := &ListNode{num % 10, nil}
		q.Next = p.Next
		p.Next = q
		num /= 10
	}
    //sum is zero
    if p.Next == nil {
        return p
    }
	return p.Next
}

func createListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	var head, p *ListNode
	for i := 0; i < len(vals); i++ {
		node := &ListNode{vals[i], nil}
		if i == 0 {
			head = node
		} else {
			p.Next = node
		}
		p = node
	}
	return head
}

func getValFromListNode(head *ListNode) []int {
	vals := make([]int, 0)
	p := head
	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}
	return vals
}

func main() {
	t1 := []int{0}
	t2 := []int{0}
	//t1 := []int{7, 2, 4, 3}
	//t2 := []int{5, 6, 4}
	l1 := createListNode(t1)
	l2 := createListNode(t2)
	t := addTwoNumbers(l1, l2)
	fmt.Println(getValFromListNode(t))
}
