package main

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := 1
	preNum := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if preNum != num {
			count--
			if count < 0 {
				preNum = num
				count = 1
			}
		} else {
			count++
		}
	}
	return preNum
}

// @lc code=end
