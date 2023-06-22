package main

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// sure to reach desination now
	count := 0
	maxIdx := nums[0] + 0
	checkPointIdx := nums[0] + 0
	for i, num := range nums {
		if checkPointIdx >= len(nums)-1 {
			count++
			break
		}
		curMaxIdx := num + i
		maxIdx = max(maxIdx, curMaxIdx)
		if checkPointIdx == i {
			count++
			checkPointIdx = maxIdx
		}
	}

	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
