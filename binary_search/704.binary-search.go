package main

/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2
		value := nums[mid]

		if value == target {
			return mid
		}
		if value > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
