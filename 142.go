package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	p, q := head, head
	for p != nil && q != nil {
		p = p.Next
		q = q.Next
		if q == nil {
			break
		}
		q = q.Next

		if p == q {
			q = head
			for p != q {
				p = p.Next
				q = q.Next
			}
			return p
		}
	}
	return nil
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

func setCycle(head *ListNode, start, end int) {
	var startNode, stopNode *ListNode
	p := head
	i := 0
	for p != nil {
		if i == start {
			startNode = p
		}
		if i == end {
			stopNode = p
		}
		p = p.Next
		i++
	}
	if startNode != nil && stopNode != nil {
		startNode.Next = stopNode
	}
}

func main() {
	head := createListNode([]int{3, 2, 0, 1, -4})
	setCycle(head, 4, 0)
	ans := detectCycle(head)
	fmt.Println(ans.Val)
}
