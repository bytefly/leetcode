package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDepths(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftAvl := getDepths(root.Left)
	rightDepth, rightAvl := getDepths(root.Right)
	if root.Left != nil {
		leftDepth++
	}
	if root.Right != nil {
		rightDepth++
	}

	depth := leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}

	if !leftAvl || !rightAvl || int(math.Abs(float64(leftDepth-rightDepth))) > 1 {
		return depth, false
	}
	return depth, true
}

func isBalanced(root *TreeNode) bool {
	_, ans := getDepths(root)
	return ans
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

			if i >= len(vals)-1 || vals[i+1] == -1 {
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
	fmt.Println(isBalanced(MakeTreeNode([]int{1, 2, 2, 3, 4, 4, 3})))
	fmt.Println(isBalanced(MakeTreeNode([]int{1, 2, 2, -1, 3, -1, 3})))
	fmt.Println(isBalanced(MakeTreeNode([]int{1, 2, 2, -1, -1, -1, 3})))
	fmt.Println(isBalanced(MakeTreeNode([]int{1})))
	fmt.Println(isBalanced(MakeTreeNode([]int{3, 9, 20, -1, -1, 15, 7})))
	fmt.Println(isBalanced(MakeTreeNode([]int{1, 2, 2, 3, 3, -1, -1, 4, 4})))
	fmt.Println(isBalanced(MakeTreeNode([]int{1, -1, 2, -1, 3})))
	fmt.Println(isBalanced(MakeTreeNode([]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, -1, -1, 5, 5})))
	fmt.Println(isBalanced(MakeTreeNode([]int{1, 2, 2, 3, 4})))
}
