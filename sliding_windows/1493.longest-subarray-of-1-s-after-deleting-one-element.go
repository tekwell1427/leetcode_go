package main

/*
 * @lc app=leetcode id=1493 lang=golang
 *
 * [1493] Longest Subarray of 1's After Deleting One Element
 */

// @lc code=start
func longestSubarray(nums []int) int {
	output := 0
	left, right := 0, 0

	for _, num := range nums {
		if num == 0 {
			left, right = right, 0
		} else {
			left++
			right++
		}
		output = max(output, left)
	}
	// if there is no zero in slice, have to minus 1
	if left == len(nums) {
		output--
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
