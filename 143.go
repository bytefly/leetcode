package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
    if head == nil {
        return
    }

	nodes := make([]*ListNode, 0)
	p := head
	for p != nil {
		nodes = append(nodes, p)
		p = p.Next
	}

	j := 0
	for i := len(nodes) - 1; i > len(nodes)/2; i-- {
		nodes[j].Next = nodes[i]
		nodes[i].Next = nodes[j+1]
		j++
	}
    nodes[len(nodes)/2].Next = nil
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
	//head := createListNode([]int{1})
	//head := createListNode([]int{1, 2})
	//head := createListNode([]int{1, 2, 3})
	//head := createListNode([]int{1, 2, 3, 4})
	head := createListNode([]int{1, 2, 3, 4, 5})
	reorderList(head)
	fmt.Println(getValFromListNode(head))
}
