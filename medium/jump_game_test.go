package medium

import "testing"

func TestJumpGame(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"len(nums) == 1, should return true", []int{1}, true},
		{"there's a way, should return true", []int{2, 3, 1, 1, 4}, true},
		{"there's no way, should return false", []int{3, 2, 1, 0, 4}, false},
		{"there's a way for [2,0]", []int{2, 0}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canJump2(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected is %v, but got %v", tc.expected, result)
			}
		})
	}
}
