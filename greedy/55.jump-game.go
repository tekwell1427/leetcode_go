package main

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	maxIdx := 0
	for i, num := range nums {
		if i > maxIdx {
			return false
		}
		curMaxIdx := i + num
		maxIdx = max(maxIdx, curMaxIdx)
		if maxIdx >= len(nums)-1 {
			return true
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
