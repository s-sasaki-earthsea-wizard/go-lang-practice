package validate_binary_search_tree

import (
	"go-lang-practice/pkg/types"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name     string
		tree     *types.TreeNode
		expected bool
	}{
		{
			name: "valid BST case",
			tree: &types.TreeNode{
				Val:   2,
				Left:  &types.TreeNode{Val: 1},
				Right: &types.TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			name: "invalid BST case",
			tree: &types.TreeNode{
				Val: 5,
				Left: &types.TreeNode{
					Val: 1,
				},
				Right: &types.TreeNode{
					Val:   4,
					Left:  &types.TreeNode{Val: 3},
					Right: &types.TreeNode{Val: 6},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidBST(tt.tree)
			t.Logf("Input tree root value: %v", tt.tree.Val)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			} else {
				t.Logf("Test passed: got %v, want %v", result, tt.expected)
			}
		})
	}
}
