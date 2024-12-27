package easy

import "unicode"

func lengthOfLastWord(s string) int {
	length := 0
	lastSymbolIsSpace := false
	for _, r := range s {
		isLetter := unicode.IsLetter(r)
		if isLetter && lastSymbolIsSpace {
			length = 0
			lastSymbolIsSpace = false
		}

		if isLetter {
			length++
		} else {
			lastSymbolIsSpace = true
		}
	}

	return length
}
