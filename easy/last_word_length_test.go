package easy

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		data     string
		expected int
	}{
		{"Hello World should return 5", "Hello World", 5},
		{"Breaking Bad with ending space should return 3", "Breaking Bad ", 3},
		{"Premier padel competition with five ending spaces should return 11", "Premier padel   competition     ", 11},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := lengthOfLastWord(tc.data)
			if result != tc.expected {
				t.Errorf("For test %s, expected result is %d but got %d", tc.name, tc.expected, result)
			}
		})
	}
}
