package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// record := make(map[byte]bool)
	record := [128]bool{}
	head := 0
	output := 0
	for i := 0; i < len(s); i++ {
		cur := s[i]
		for record[cur] {
			record[s[head]] = false
			head++
		}
		record[cur] = true
		output = max(output, i-head+1)
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
