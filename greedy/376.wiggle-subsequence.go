package main

/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		// 0 or 1
		return len(nums)
	}
	preDiff := 0
	output := 1

	for i := 0; i < len(nums)-1; i++ {
		curDiff := nums[i+1] - nums[i]

		if preDiff <= 0 && curDiff > 0 || preDiff >= 0 && curDiff < 0 {
			output++
			preDiff = curDiff
		}
	}
	return output
}

// @lc code=end
