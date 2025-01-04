package main

import "fmt"

// TreeNode defines the structure of a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Calculate the maximum depth of a binary tree
func maxDepth(root *TreeNode) int {
	// If the root is nil, return 0
	if root == nil {
		return 0
	}

	// Calculate the maximum depth of the left and right subtrees
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// Return the maximum depth of the left and right subtrees plus 1
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	// Create a binary tree
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	result := maxDepth(root)
	fmt.Println(result)
}
