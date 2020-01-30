package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublits(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0)
	p := head
	var sum, curSum int
	var find bool

	for p != nil {
		if p.Val == 0 {
			p = p.Next
			continue
		}

		find = false
		curSum = p.Val
		for j := len(nodes) - 1; j >= 0; j-- {
			curSum += nodes[j].Val
			if curSum == 0 {
				nodes = nodes[:j]
				find = true
				sum -= (curSum - p.Val)
				break
			}
		}

		if !find {
			sum += p.Val
			nodes = append(nodes, p)
		}
		p = p.Next
	}

	if len(nodes) == 0 {
		return nil
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
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
	p := removeZeroSumSublits(createListNode([]int{1, 2, 3, -4, 1, 5}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{0, 2, -2, -2, 1, 2, 2, -1, -2}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{1, 2, -3, 3, 1}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{1, 2, 3, -3, 4}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{1, 2, 3, -3, -2}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{0, 2, 3, -3, -2}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{2}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{0}))
	fmt.Println(getValFromListNode(p))
	p = removeZeroSumSublits(createListNode([]int{}))
	fmt.Println(getValFromListNode(p))
}
