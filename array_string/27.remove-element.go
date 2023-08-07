package main

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	curInsertIdx := 0
	for _, num := range nums {
		if num == val {
			continue
		}
		nums[curInsertIdx] = num
		curInsertIdx++
	}
	return curInsertIdx
}

// @lc code=end
