package easy

import "testing"

func Test1(t *testing.T) {
	m := 4
	n := 3
	nums1 := []int{1, 3, 8, 12, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, m, nums2, n)

	result := []int{1, 2, 3, 5, 6, 8, 12}
	for i := 0; i < 7; i++ {
		if result[i] != nums1[i] {
			t.Errorf("Wrong answer, has to be %v, but got %v", result, nums1)
			break
		}
	}
}
