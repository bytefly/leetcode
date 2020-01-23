package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	q := head
	var newHead, dummy *ListNode
	same := false

	for q != nil {
		if q.Next == nil || q.Val != q.Next.Val {
			if dummy != nil {
				dummy.Next = q
			} else {
				newHead = q
			}
			dummy = q

			if same {
				same = false
				//if end with same node,drop others
				if q.Next == nil && dummy != nil {
					dummy.Next = nil
				}
			}
		} else if q.Next != nil && q.Val == q.Next.Val {
			same = true
		}

		q = q.Next
	}

	return newHead
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
	head := createListNode([]int{1, 2, 2, 3, 4, 4, 5})
	//head := createListNode([]int{2, 2})
	//head := createListNode([]int{1, 2, 2})
	//head := createListNode([]int{2, 2, 3})
	//head := createListNode([]int{2, 2, 2})
	//head := createListNode([]int{1, 2, 2, 3})
	//head := createListNode([]int{2})
	//head := createListNode([]int{2, 1})
	head = deleteDuplicates(head)

	vals := getValFromListNode(head)
	fmt.Println(vals)
}
