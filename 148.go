package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

    //move to the middle node
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
        fast = fast.Next
	}

    //get left and right ordered list recursively
	right := slow.Next
	slow.Next = nil
	l, r := sortList(head), sortList(right)

    //merge left and right ordered list together
	dummy := &ListNode{0, nil}
	head = dummy
	for l != nil && r != nil {
		if l.Val <= r.Val {
			head.Next = l
			head, l = head.Next, l.Next
		} else {
			head.Next = r
			head, r = head.Next, r.Next
		}
	}
	if l != nil {
		head.Next = l
	}
	if r != nil {
		head.Next = r
	}
	return dummy.Next
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
	//head := createListNode([]int{3, 4, 1})
	head := createListNode([]int{-1, 5, 3, 4, 0})
	//head := createListNode([]int{4, 2, 1, 3})
	//head := createListNode([]int{3, 1})
	//head := createListNode([]int{3})
	head = sortList(head)
	fmt.Println(getValFromListNode(head))
}
