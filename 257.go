package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreePath struct {
	Node *TreeNode
	Path string
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var t *TreePath
	ans := make([]string, 0)
	rootPath := &TreePath{root, strconv.FormatInt(int64(root.Val), 10)}
	stack := make([]*TreePath, 0)
	stack = append(stack, rootPath)

	for len(stack) > 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			p := stack[i]
			if p.Node.Left == nil && p.Node.Right == nil {
				ans = append(ans, p.Path)
			}
			if p.Node.Left != nil {
				t = &TreePath{p.Node.Left, p.Path + "->" + strconv.FormatInt(int64(p.Node.Left.Val), 10)}
				stack = append(stack, t)
			}
			if p.Node.Right != nil {
				t = &TreePath{p.Node.Right, p.Path + "->" + strconv.FormatInt(int64(p.Node.Right.Val), 10)}
				stack = append(stack, t)
			}
		}
		stack = stack[size:]
	}

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
	root := MakeTreeNode([]int{1, 2, 3, -1, 4, -1, 5})
	fmt.Println(binaryTreePaths(root))
	root = MakeTreeNode([]int{1, 2, 3, -1, -1, -1, 4})
	fmt.Println(binaryTreePaths(root))
	root = MakeTreeNode([]int{1})
	fmt.Println(binaryTreePaths(root))
	root = MakeTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(binaryTreePaths(root))
	root = MakeTreeNode([]int{1, -1, 2, -1, 3})
	fmt.Println(binaryTreePaths(root))
}
