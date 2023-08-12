package main

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	curJumpMaxIdx := 0

	for i, num := range nums {
		if i > curJumpMaxIdx {
			return false
		}
		curJumpMaxIdx = max(curJumpMaxIdx, i+num)
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
