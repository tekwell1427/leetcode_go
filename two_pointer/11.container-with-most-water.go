package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	output := 0
	for left < right {
		l := height[left]
		r := height[right]
		output = max(output, (right-left)*min(l, r))
		if l > r {
			right--
		} else {
			left++
		}
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
