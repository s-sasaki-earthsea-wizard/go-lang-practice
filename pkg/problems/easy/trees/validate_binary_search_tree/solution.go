package validate_binary_search_tree

import "go-lang-practice/pkg/types"

// Validate if a binary tree is a binary search tree
func isValidBST(root *types.TreeNode) bool {
	return validateBST(root, nil, nil)
}

func validateBST(node *types.TreeNode, min, max *int) bool {
	// If the root is nil, return true
	if node == nil {
		return true
	}

	// If there is a minimum value and the root value is less than the minimum value, return false
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	// Recursively validate the left and right subtrees
	return validateBST(node.Left, min, &node.Val) &&
		validateBST(node.Right, &node.Val, max)
}
