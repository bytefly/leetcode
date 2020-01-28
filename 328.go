package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEventList(head *ListNode) *ListNode {
    i := 0
    odd, event := &ListNode{0, nil}, &ListNode{0, nil}
    eventHead := event

    p := head
    for p != nil {
        if i % 2 == 0 {
            odd.Next = p
            odd = odd.Next
        } else {
            event.Next = p
            event = event.Next
        }

        i++
        p=p.Next
    }
    event.Next = nil
    odd.Next = eventHead.Next

    return head
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
	//t := []int{1}
	//t := []int{1, 2}
	//t := []int{1, 2, 3}
	//t := []int{4, 5, 1, 9}
	t := []int{1, 2, 3, 4, 5}
	head := createListNode(t)
    head = oddEventList(head)
	fmt.Println(getValFromListNode(head))
}
