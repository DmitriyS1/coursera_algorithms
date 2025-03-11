package medium

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		name     string
		articles []int
		expected int
	}{
		{name: "empty articles", articles: []int{}, expected: 0},
		{name: "single article", articles: []int{1}, expected: 1},
		{name: "single article with more than 1 articles", articles: []int{22}, expected: 1},

		{name: "two articles", articles: []int{1, 2}, expected: 1},
		{name: "two articles reversed", articles: []int{3, 1}, expected: 1},
		{name: "three articles", articles: []int{2, 1, 3}, expected: 2},
		{name: "three articles in wrong order", articles: []int{1, 3, 1}, expected: 1},
		{name: "three articles without citations", articles: []int{0, 0, 2}, expected: 1},
		{name: "four articles", articles: []int{3, 4, 1, 2}, expected: 2},
		{name: "six articles", articles: []int{2, 4, 8, 9, 9, 3}, expected: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := hIndex(test.articles)
			if result != test.expected {
				t.Errorf("expected: %d, got: %d", test.expected, result)
				return
			}
		})
	}
}
