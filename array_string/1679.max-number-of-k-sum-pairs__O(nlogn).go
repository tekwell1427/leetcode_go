package main

import "sort"

/*
 * @lc app=leetcode id=1679 lang=golang
 *
 * [1679] Max Number of K-Sum Pairs
 */

// @lc code=start
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	output := 0

	left, right := 0, len(nums)-1
	for left < right {
		s := nums[left] + nums[right]
		if s == k {
			output++
			left++
			right--
		} else if s < k {
			left++
		} else {
			right--
		}
	}
	return output
}

// @lc code=end
