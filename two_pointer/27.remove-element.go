package main

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	slow := 0

	for _, num := range nums {
		if num != val {
			nums[slow] = num
			slow++
		}
	}
	return slow
}

// @lc code=end
