package find_middle

import "testing"

func buildList(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := &Node{Val: vals[0]}
	current := head
	for _, v := range vals[1:] {
		current.Next = &Node{Val: v}
		current = current.Next
	}
	return head
}

func TestFindMiddle(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected *Node
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "one element",
			input:    []int{10},
			expected: &Node{Val: 10},
		},
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: &Node{Val: 2},
		},
		{
			name:     "odd number of elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: &Node{Val: 3},
		},
		{
			name:     "even number of elements",
			input:    []int{1, 2, 3, 4},
			expected: &Node{Val: 3},
		},
		{
			name:     "long list",
			input:    []int{5, 7, 9, 11, 13, 15, 17},
			expected: &Node{Val: 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)

			mid := FindMiddle(head)

			if tt.expected == nil {
				if mid != nil {
					t.Errorf("expected nil, got %v", mid)
				}
				return
			}

			if mid == nil || mid.Val != tt.expected.Val {
				t.Errorf("expected middle %v, got %v", tt.expected, mid)
			}
		})
	}
}
