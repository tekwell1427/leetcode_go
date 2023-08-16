package main

/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1

	curIdx := len(nums) - 1
	output := make([]int, len(nums))
	for left < right {
		l := sqr(nums[left])
		r := sqr(nums[right])
		if l > r {
			output[curIdx] = l
			left++
		} else {
			output[curIdx] = r
			right--
		}
		curIdx--
	}
	output[curIdx] = sqr(nums[left])
	return output
}
func sqr(a int) int {
	return a * a
}

// @lc code=end
