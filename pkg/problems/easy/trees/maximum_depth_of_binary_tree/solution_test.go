package max_depth_of_binary_tree

import (
	"go-lang-practice/pkg/types"
	"testing"
)

// TreeNode defines the structure of a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestMaxDepth(t *testing.T) {
	// Create a binary tree
	root := &types.TreeNode{
		Val:  3,
		Left: &types.TreeNode{Val: 9},
		Right: &types.TreeNode{
			Val:   20,
			Left:  &types.TreeNode{Val: 15},
			Right: &types.TreeNode{Val: 7},
		},
	}
	result := MaxDepth(root)
	expected := 3

	// Test the result
	if result != expected {
		t.Errorf("Expected 3, but got %v", result)
	} else {
		t.Logf("Success! Expected: %v, Got: %v", expected, result)
	}
}
