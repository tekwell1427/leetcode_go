package main

import "math"

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	first, second := nums[0], math.MaxInt32
	for _, num := range nums {
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}

// @lc code=end
