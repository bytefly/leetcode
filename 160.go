package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var ans *ListNode
	tail := headB
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}
	//make a ring
	tail.Next = headB

	//check if the cycle exists
	slow, fast := headA, headA
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if slow == fast {
			break
		}
	}
	//find the intersection node
	if slow != nil && slow == fast {
		fast = headA
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		ans = slow
	}

	//disconnect the ring
	tail.Next = nil
	return ans
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
	/*
		headA := createListNode([]int{1})
		headB := createListNode([]int{5, 0, 1})
	*/
	headA := createListNode([]int{4, 1})
	headB := createListNode([]int{5, 0, 1})
	headC := createListNode([]int{8, 4, 5})
	p := headA
	for p != nil && p.Next != nil {
		p = p.Next
	}
	p.Next = headC

	p = headB
	for p != nil && p.Next != nil {
		p = p.Next
	}
	p.Next = headC

	ans := getIntersectionNode(headA, headB)
	if ans != nil {
		fmt.Println(ans.Val)
	} else {
		fmt.Println("nil")
	}
}
