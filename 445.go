package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := make([]int, 0), make([]int, 0)
	p := l1
	for p != nil {
		num1 = append(num1, p.Val)
		p = p.Next
	}

	p = l2
	for p != nil {
		num2 = append(num2, p.Val)
		p = p.Next
	}

	p = &ListNode{0, nil}
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; carry > 0 || i >= 0 || j >= 0; i, j = i-1, j-1 {
		q := &ListNode{0, nil}
		a, b := 0, 0
		switch {
		case i >= 0 && j >= 0:
			a, b = num1[i], num2[j]
		case i < 0 && j >= 0:
			b = num2[j]
		case i >= 0 && j < 0:
			a = num1[i]
		}

		q.Val = a + b + carry
		if q.Val >= 10 {
			carry = 1
			q.Val -= 10
		} else {
            carry = 0
        }

		q.Next = p.Next
		p.Next = q
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
	t1 := []int{1}
	t2 := []int{9}
	//t1 := []int{7, 2, 4, 3}
	//t2 := []int{5, 6, 4}
	l1 := createListNode(t1)
	l2 := createListNode(t2)
	t := addTwoNumbers(l1, l2)
	fmt.Println(getValFromListNode(t))
}
