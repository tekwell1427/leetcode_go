package main

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	count, cur := 1, 1
	preNum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == preNum {
			count++
			if count > 2 {
				continue
			}
		} else {
			count = 1
			preNum = nums[i]
		}
		nums[cur] = nums[i]
		cur++
	}
	return cur
}

// @lc code=end
