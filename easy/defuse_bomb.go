package easy

/*
You have a bomb to defuse, and your time is running out! Your informer will provide you with a circular array code of length of n and a key k.

To decrypt the code, you must replace every number. All the numbers are replaced simultaneously.

    If k > 0, replace the ith number with the sum of the next k numbers.
    If k < 0, replace the ith number with the sum of the previous k numbers.
    If k == 0, replace the ith number with 0.

As code is circular, the next element of code[n-1] is code[0], and the previous element of code[0] is code[n-1].

Given the circular array code and an integer key k, return the decrypted code to defuse the bomb!


Example 1:

Input: code = [5,7,1,4], k = 3
Output: [12,10,16,13]
Explanation: Each number is replaced by the sum of the next 3 numbers. The decrypted code is [7+1+4, 1+4+5, 4+5+7, 5+7+1]. Notice that the numbers wrap around.

Example 2:

Input: code = [1,2,3,4], k = 0
Output: [0,0,0,0]
Explanation: When k is zero, the numbers are replaced by 0.

Example 3:

Input: code = [2,4,9,3], k = -2
Output: [12,5,6,13]
Explanation: The decrypted code is [3+9, 2+3, 4+2, 9+4]. Notice that the numbers wrap around again. If k is negative, the sum is of the previous numbers.


Constraints:

    n == code.length
    1 <= n <= 100
    1 <= code[i] <= 100
    -(n - 1) <= k <= n - 1
*/

func decrypt(code []int, k int) []int {
	length := len(code)
	if k == 0 {
		return make([]int, length)
	}

	mirror := make([]int, 2*length-1)
	center := len(mirror) / 2
	for i := 0; i < len(mirror); i++ {
		if i < center {
			mirror[i] = code[length-1-i]
		} else {
			mirror[i] = code[i-center]
		}
	}

	result := make([]int, length)
	j := 0
	for i := center; i < len(mirror); i++ {
		if k < 0 {
			for p := k; p > 0; p++ {
				result[j] += mirror[i+p]
				j++
			}
		} else {
			for p := k; p > 0; p-- {
				if i+p >= len(mirror) {
					result[j] += mirror[length-p+center]
				}
			}
		}
	}

	return mirror
}

func decrypt2(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}

	result := make([]int, len(code))
	for i := 0; i < len(result); i++ {
		if k > 0 {
			start := 0
			for j := k; j > 0; j-- {
				lenToEnd := len(code) - i
				if i+j > len(code)-1 {
					result[i] += code[start+j-lenToEnd]
				} else {
					result[i] += code[i+j]
				}
			}
		} else {
			end := len(code)
			for j := k; j < 0; j++ {
				if i < -j {
					result[i] += code[end+j+i]
				} else {
					result[i] += code[j+i]
				}
			}
		}
	}

	return result
}
