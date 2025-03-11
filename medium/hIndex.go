package medium

import "sort"

func hIndex(citations []int) int {
	hIndex := 0

	cLen := len(citations)
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		cits := 0
		if citations[i] > cLen {
			cits = cLen
		} else {
			cits = citations[i]
		}

		if cits > hIndex && cLen >= cits {
			hIndex = cits
		}
		cLen--
	}

	return hIndex
}
