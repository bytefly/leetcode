package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	MinVals []int
	Index   int
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{make([]int, 0), 0}
	stack := make([]*TreeNode, 0)
	p := root
	for {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) > 0 {
			p = stack[len(stack)-1]
			iterator.MinVals = append(iterator.MinVals, p.Val)
			stack = stack[:len(stack)-1]
			p = p.Right
		} else {
			break
		}
	}

	return iterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.Index >= len(this.MinVals) {
		return -1
	}

	min := this.MinVals[this.Index]
	this.Index++
	return min
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.Index >= len(this.MinVals) {
		return false
	}
	return true
}

func MakeTreeNode(vals []int) *TreeNode {
	var root, node0, node1, node2 *TreeNode
	stack := make([]*TreeNode, 0)

	for i := 0; i < len(vals); {
		if len(stack) == 0 {
			node0 = &TreeNode{vals[i], nil, nil}
			if i == 0 {
				root = node0
			}
			stack = append(stack, node0)
			i++
		} else {
			node1, node2 = nil, nil
			node0 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if vals[i] != -1 {
				node1 = &TreeNode{vals[i], nil, nil}
				node0.Left = node1
			}
			if vals[i+1] != -1 {
				node2 = &TreeNode{vals[i+1], nil, nil}
				node0.Right = node2
			}

			if node2 != nil {
				stack = append(stack, node2)
			}
			if node1 != nil {
				stack = append(stack, node1)
			}

			i += 2
		}
	}

	return root
}

func main() {
	root := MakeTreeNode([]int{7, 3, 15, -1, -1, 9, 20})
	obj := Constructor(root)
	fmt.Println(obj.Next())
	fmt.Println(obj.Next())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
}
