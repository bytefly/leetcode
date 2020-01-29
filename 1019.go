package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
    nums := make([]int, 0)
    monoStack := make([]int, 0)
	p := head
	for p != nil {
        nums = append(nums, p.Val)
		p = p.Next
	}

	ans := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        for len(monoStack) > 0 && nums[monoStack[len(monoStack)-1]] < nums[i] {
            ans[monoStack[len(monoStack)-1]] = nums[i]
            monoStack = monoStack[:len(monoStack)-1]
        }
        monoStack = append(monoStack, i)
    }
	return ans
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
	p := nextLargerNodes(createListNode([]int{2, 1, 5}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{2, 7, 4, 3, 5}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{1, 7, 5, 1, 9, 2, 5, 1}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{2}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{}))
	fmt.Println(p)
}
