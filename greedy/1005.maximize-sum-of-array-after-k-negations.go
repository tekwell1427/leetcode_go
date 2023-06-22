package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=1005 lang=golang
 *
 * [1005] Maximize Sum Of Array After K Negations
 */

// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	sum := 0
	minNum := 101
	for _, num := range nums {
		if num < 0 && k > 0 {
			k--
			num = (-num)
		}
		minNum = min(minNum, num)
		sum += num
	}
	if k%2 == 1 {
		sum = sum - 2*minNum
	}
	return sum
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
