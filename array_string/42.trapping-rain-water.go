package main

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	curMax := 0
	for i := 0; i < len(height); i++ {
		curMax = max(curMax, height[i])
		leftMax[i] = curMax
	}
	curMax = 0
	for i := len(height) - 1; i >= 0; i-- {
		curMax = max(curMax, height[i])
		rightMax[i] = curMax
	}
	output := 0
	for i := 0; i < len(height); i++ {
		m := min(leftMax[i], rightMax[i])
		output += m - height[i]
	}
	return output
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
