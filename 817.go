package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	var find bool
    var q *ListNode

	p, num := head, 0
	for p != nil {
		find = false
		for j := 0; j < len(G); j++ {
			if p.Val == G[j] {
				find = true
				break
			}
		}

		if !find {
            //only the bound counts
			if q != nil && q.Next == p {
				num++
			}
		} else {//save the last node in G
            q = p
        }
		p = p.Next
	}
    //handle the last many nodes make a plugin
	if find {
		num++
	}
	return num
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
	fmt.Println(numComponents(nil, []int{1}))
	fmt.Println(numComponents(createListNode([]int{0}), []int{0}))
	fmt.Println(numComponents(createListNode([]int{3, 4, 0, 2, 1}), []int{4}))
	fmt.Println(numComponents(createListNode([]int{0, 1, 2}), []int{0, 2}))
	fmt.Println(numComponents(createListNode([]int{0, 1, 2, 3}), []int{0, 1, 3}))
	fmt.Println(numComponents(createListNode([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}))
}
