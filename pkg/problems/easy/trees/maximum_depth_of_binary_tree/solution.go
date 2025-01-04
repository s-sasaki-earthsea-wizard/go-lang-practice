package max_depth_of_binary_tree

import "go-lang-practice/pkg/types"

// Calculate the maximum depth of a binary tree
func MaxDepth(root *types.TreeNode) int {
	// If the root is nil, return 0
	if root == nil {
		return 0
	}

	// Calculate the maximum depth of the left and right subtrees
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	// Return the maximum depth of the left and right subtrees plus 1
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
