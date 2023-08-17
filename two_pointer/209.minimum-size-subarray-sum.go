package main

import (
	"math"
)

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	left := 0
	curSum, minSubLen := 0, math.MaxInt32

	for right := 0; right < len(nums); right++ {
		curSum += nums[right]
		for curSum >= target {
			minSubLen = min(minSubLen, right-left+1)
			curSum -= nums[left]
			left++
		}
	}
	if minSubLen == math.MaxInt32 {
		return 0
	}
	return minSubLen

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
