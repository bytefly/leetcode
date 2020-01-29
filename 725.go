package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	var length, partSize int
	p := root
	for p != nil {
		p = p.Next
		length++
	}

	partLen := length / k
	remainSize := length % k
	parts := make([]*ListNode, k, k)
	p = root
	for i := 0; i < k; i++ {
		switch {
		case partLen == 0 && i < remainSize:
			partSize = 1
		case partLen == 0 && i >= remainSize:
			partSize = 0
		case partLen > 0 && i < remainSize:
			partSize = 1 + partLen
		case partLen > 0 && i >= remainSize:
			partSize = partLen
		}

		if partSize == 0 {
			parts[i] = nil
			continue
		}

		q := parts[i]
		for j := 0; j < partSize; j++ {
			if q == nil {
				parts[i] = p
				q = p
			} else {
				q.Next = p
				q = p
			}
			p = p.Next
		}
		q.Next = nil
	}

	return parts
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
	vals := []int{1, 2, 3, 4}
	t := splitListToParts(createListNode(vals), 5)
	//vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//t := splitListToParts(createListNode(vals), 3)
	for i := 0; i < len(t); i++ {
		fmt.Println(getValFromListNode(t[i]))
	}
}
