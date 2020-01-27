package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	p, q := head, dummy
	for p != nil {
		if p.Val != val {
			q.Next = p
			q = p
		}
		p = p.Next
	}
    //handle the last node removed
    q.Next = nil
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
	t := []int{6, 2, 6, 3, 4, 5, 6}
	head := createListNode(t)
	head = removeElements(head, 6)
	fmt.Println(getValFromListNode(head))
}
