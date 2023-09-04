package main

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	record := make(map[int]bool, len(nums))

	for _, num := range nums {
		record[num] = true
	}
	output := 0
	for key := range record {
		if _, exists := record[key-1]; !exists {
			length := 0
			for record[key] {
				length++
				key++
			}
			output = max(output, length)
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
