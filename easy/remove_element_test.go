package easy

import "testing"

func TestRemoveElement(t *testing.T) {
	tCases := []struct {
		name     string
		input    []int
		val      int
		expected int
	}{
		{name: "array with one element != to val => expected 1", input: []int{1}, val: 2, expected: 1},
		{name: "array with 4 elements containing val => 2", input: []int{3, 2, 2, 3}, val: 3, expected: 2},
		{name: "array with 8 elements containing val => 5", input: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expected: 5},
		{name: "failed case from leetCode", input: []int{4, 2, 0, 2, 2, 1, 4, 4, 1, 4, 3, 2}, val: 4, expected: 8},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeElement(tc.input, tc.val)
			if res != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, res)
			}
		})
	}
}
