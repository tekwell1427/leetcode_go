package main

/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

// @lc code=start
func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for i, num := range nums {
		sum = sum - num
		if sum == leftSum {
			return i
		}
		leftSum += num
	}
	return -1
}

// @lc code=end
