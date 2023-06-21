package main

/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := 1
	preDiff := nums[1] - nums[0]
	if preDiff != 0 {
		count = 2
	}
	for i := 2; i < len(nums); i++ {
		curDiff := (nums[i] - nums[i-1])
		if (preDiff >= 0 && curDiff < 0) || (preDiff <= 0 && curDiff > 0) {
			count++
			// third situation:
			// only change positive/negtive vector when we meet peak/valley
			preDiff = curDiff
		}
		// preDiff = curDiff
	}
	return count

}

// @lc code=end
