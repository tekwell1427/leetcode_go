package main

/*
 * @lc app=leetcode id=1456 lang=golang
 *
 * [1456] Maximum Number of Vowels in a Substring of Given Length
 */

// @lc code=start
func maxVowels(s string, k int) int {
	vowels := make(map[byte]bool)
	for _, vowel := range []byte{'a', 'e', 'i', 'o', 'u'} {
		vowels[vowel] = true
	}
	output := 0

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			output++
		}
	}
	count := output
	left := 0
	for i := k; i < len(s); i++ {

		if vowels[s[left]] {
			count--
		}
		left++

		if vowels[s[i]] {
			count++
		}
		output = max(output, count)
	}
	return output
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
