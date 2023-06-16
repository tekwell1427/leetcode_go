package main

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	output := [][]string{}
	var helper func(preIdx int, subString [][]byte)
	helper = func(preIdx int, subString [][]byte) {
		if preIdx == len(s)-1 {
			tmp := make([]string, len(subString))
			for i := 0; i < len(subString); i++ {
				tmp[i] = string(subString[i])
			}
			output = append(output, tmp)
			return
		}
		curS := []byte{}
		for i := preIdx + 1; i < len(s); i++ {
			character := s[i]
			curS = append(curS, character)
			if isPalindrone(curS) {
				subString = append(subString, curS)
				helper(i, subString)
				subString = subString[:len(subString)-1]
			}
		}
	}
	helper(-1, [][]byte{})
	return output
}
func isPalindrone(a []byte) bool {
	left, right := 0, len(a)-1
	for left < right {
		if a[left] != a[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end
