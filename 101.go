package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	vals := make([]int, 0)
	for len(stack) > 0 {
		size := len(stack)
		halfNodeLen, leafNum := size/2, 0
		for i := 0; i < size; i++ {
			p, val := stack[i], -1
			if p != nil {
				val = p.Val
				stack = append(stack, p.Left)
				stack = append(stack, p.Right)
			} else {
				stack = append(stack, nil)
				stack = append(stack, nil)
				leafNum++
			}

			//compare with the mirror node in the front half nodes
			if size > 1 && i >= halfNodeLen {
				if val != vals[halfNodeLen*2-i-1] {
					return false
				}
			} else {
				vals = append(vals, val)
			}
		}
		//stop if no subtree exists
		if leafNum == size {
			break
		}
		//clear the half node values
		vals = vals[:0]
		stack = stack[size:]
	}

	return true
}

func GetTreeNodeVal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	vals := make([]int, 0)
	for len(stack) > 0 {
		size := len(stack)
		nilNum := 0
		for i := 0; i < size; i++ {
			p := stack[i]
			if p != nil {
				vals = append(vals, p.Val)
				stack = append(stack, p.Left)
				stack = append(stack, p.Right)
			} else {
				vals = append(vals, -1)
				nilNum++
			}
		}
		//trim all nil leaf nodes
		if nilNum == size {
			vals = vals[:len(vals)-size]
		}
		stack = stack[size:]
	}
	return vals
}

func MakeTreeNode(vals []int) *TreeNode {
	var root, node0, node1, node2 *TreeNode
	queue := make([]*TreeNode, 0)
	i := 0

	for {
		if i == 0 {
			if i >= len(vals) {
				return root
			}
			node0 = &TreeNode{vals[i], nil, nil}
			root = node0
			queue = append(queue, node0)
			i++
			continue
		}

		node0Num := len(queue)
		for j := 0; j < node0Num; j++ {
			node0 = queue[j]

			if i >= len(vals) {
				return root
			}
			if node0 == nil {
				//no children parse
				continue
			}

			if vals[i] == -1 {
				node1 = nil
			} else {
				node1 = &TreeNode{vals[i], nil, nil}
			}
			node0.Left = node1
			queue = append(queue, node1)

			if vals[i+1] == -1 {
				node2 = nil
			} else {
				node2 = &TreeNode{vals[i+1], nil, nil}
			}
			node0.Right = node2
			queue = append(queue, node2)

			i += 2
		}
		queue = queue[node0Num:]
	}

	return root
}

func main() {
	root := MakeTreeNode([]int{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(isSymmetric(root))
	root = MakeTreeNode([]int{1, 2, 2, -1, 3, -1, 3})
	fmt.Println(isSymmetric(root))

	root = MakeTreeNode([]int{9, -42, -42, -1, 76, 76, -1, -1, 13, -1, 13})
	fmt.Println(GetTreeNodeVal(root))
	fmt.Println(GetTreeNodeVal(MakeTreeNode([]int{1, 2, 3, -1, 4, 5, -1})))
	fmt.Println(isSymmetric(root))
}
