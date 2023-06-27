package main

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
func partitionLabels(s string) []int {
	record := make(map[rune]int)
	// record last appearing index
	for i, char := range s {
		record[char] = i
	}
	output := []int{}
	maxIdx := 0
	headIdx := -1
	for i, char := range s {
		maxIdx = max(maxIdx, record[char])
		if maxIdx == i {
			output = append(output, maxIdx-headIdx)
			headIdx = maxIdx
		}
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
