package main

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	var output int
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2
		midNum := nums[mid]
		if midNum == target {
			return mid
		} else if midNum < target {
			left = mid + 1
		} else if midNum > target {
			right = mid
		}
	}
	if left < len(nums) {
		if nums[left] > target {
			output = left
		} else {
			output = left + 1
		}
	} else {
		output = left
	}
	return output
}
func abs(a, b int) int {
	output := a - b
	if output < 0 {
		output = -output
	}
	return output
}

// @lc code=end
