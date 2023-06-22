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
	maxRangeIdx := nums[0] + 0
	checkPointIdx := nums[0] + 0
	for i := 1; i < len(nums); i++ {
		curMaxIdx := nums[i] + i
		maxRangeIdx = max(maxRangeIdx, curMaxIdx)
		if checkPointIdx >= len(nums)-1 {
			count++
			break
		}
		if checkPointIdx == i {
			count++
			checkPointIdx = maxRangeIdx
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
