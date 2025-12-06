package mergelists

import (
	"reflect"
	"testing"
)

// BuildList строит связный список из массива чисел.
func BuildList(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := &Node{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &Node{Val: v}
		curr = curr.Next
	}
	return head
}

// ListToSlice преобразует список в срез (для проверки в тестах).
func ListToSlice(head *Node) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "both lists empty",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "first list empty",
			list1:    []int{},
			list2:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "second list empty",
			list1:    []int{1, 2, 3},
			list2:    []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "merging two sorted lists",
			list1:    []int{1, 3, 5},
			list2:    []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "lists with duplicates",
			list1:    []int{1, 2, 2},
			list2:    []int{2, 3},
			expected: []int{1, 2, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := BuildList(tt.list1)
			l2 := BuildList(tt.list2)

			result := MergeLists(l1, l2)
			got := ListToSlice(result)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
