package main

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	preNum := nums[0]
	cur := 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if preNum != num {
			nums[cur] = num
			cur++
			preNum = num
		}
	}
	// fmt.Println(nums)
	return cur
}

// @lc code=end
