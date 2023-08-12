package main

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	curJumpMaxIdx := 0
	nextJumMaxIdx := 0
	numJumps := 0

	for i := 0; i < len(nums); i++ {
		nextJumMaxIdx = max(nextJumMaxIdx, i+nums[i])
		if i == curJumpMaxIdx && i != len(nums)-1 {
			numJumps++
			curJumpMaxIdx = nextJumMaxIdx
			nextJumMaxIdx = 0
		}
	}
	return numJumps
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
