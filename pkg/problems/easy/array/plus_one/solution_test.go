package plus_one

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			name:     "regular case",
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "with carry over",
			digits:   []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PlusOne(tt.digits)
			t.Logf("Input: digits=%v", tt.digits)
			t.Logf("Output: %v", result)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			} else {
				t.Logf("Test passed with result: %v", result)
			}
		})
	}
}
