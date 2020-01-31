package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var index int
	p, q := head, head
	dummy := &ListNode{0, nil}
	tmp := dummy

	if k <= 0 {
		return head
	}
	for q != nil {
		if index < k {
			//move the front pointer
			q = q.Next
			index++
		}
		if index == k {
			//reverse the segment
			m := p
			for {
				t := p.Next
				p.Next = tmp.Next
				tmp.Next = p
				p = t

				if p == q {
					break
				}
			}
			//reset index and tail pointer
			tmp = m
			index = 0
		}
	}

	//attach the remainder
	if index < k {
		tmp.Next = p
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
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{1, 2, 3, 4, 5}), 0)))
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{1, 2, 3, 4, 5}), 1)))
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{1, 2, 3, 4, 5}), 2)))
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{1, 2, 3, 4, 5}), 3)))
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{1, 2, 3, 4, 5}), 4)))
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{1, 2, 3, 4, 5}), 5)))
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{1, 2, 3, 4, 5}), 6)))
	fmt.Println(getValFromListNode(reverseKGroup(createListNode([]int{}), 2)))
}
