package intersection02

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "basic test",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersect(tt.nums1, tt.nums2)
			t.Logf("Input: nums1=%v, nums2=%v", tt.nums1, tt.nums2)

			// Sort the result and expected before comparing them
			sort.Ints(result)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			} else {
				t.Logf("Test passed with sorted result: %v", result)
			}
		})
	}
}
