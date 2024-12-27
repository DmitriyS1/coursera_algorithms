package easy

import "testing"

func TestMirroring(t *testing.T) {
	data := []int{3, 5, 6, 7, 8}

	expected := []int{8, 7, 6, 5, 3, 5, 6, 7, 8}

	result := decrypt(data, 12)

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Wrong result. Expected: %v, Got %v", expected, result)
		}
	}
}

func TestDecrypt2(t *testing.T) {
	testCases := []struct {
		name     string
		code     []int
		k        int
		expected []int
	}{
		{"k=0 => expected filled with zeros", []int{4, 5, 66, 75, 90}, 0, []int{0, 0, 0, 0, 0}},
		{"k=2 => expected filled with sum of two trailing numbers", []int{2, 5, 6, 8, 9}, 2, []int{11, 14, 17, 11, 7}},
		{"k=-3 => expected filled with sum of tree previous numbers", []int{2, 5, 6, 7, 11, 13}, -3, []int{31, 26, 20, 13, 18, 24}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := decrypt2(tc.code, tc.k)
			for i := 0; i < len(result); i++ {
				if result[i] != tc.expected[i] {
					t.Errorf("For test %s, expected result is %v, got %v", tc.name, tc.expected, result)
					break
				}
			}
		})
	}

}
