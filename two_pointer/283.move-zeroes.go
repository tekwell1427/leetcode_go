package main

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	curIdx := 0
	for _, num := range nums {
		if num != 0 {
			nums[curIdx] = num
			curIdx++
		}
	}
	for i := curIdx; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end
