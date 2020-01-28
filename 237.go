package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
    p:= node.Next
    if p == nil {
        return
    }
    node.Val = p.Val
    node.Next = p.Next
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

func getSliceFromListNode(head *ListNode) []*ListNode{
	nodes := make([]*ListNode, 0)
	p := head
	for p != nil {
        nodes = append(nodes, p)
		p = p.Next
	}
	return nodes
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
	t := []int{4, 5, 1, 9}
	head := createListNode(t)
    nodes := getSliceFromListNode(head)
    deleteNode(nodes[2])
	fmt.Println(getValFromListNode(head))
}
