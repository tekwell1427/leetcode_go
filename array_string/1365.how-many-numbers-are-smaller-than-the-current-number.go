package main

/*
 * @lc app=leetcode id=1365 lang=golang
 *
 * [1365] How Many Numbers Are Smaller Than the Current Number
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	record := [101]int{}

	for _, num := range nums {
		record[num]++
	}
	curSum := 0
	for i, num := range record {
		record[i] = curSum
		curSum += num
	}
	for i, num := range nums {
		// use same slice, to save meemory
		nums[i] = record[num]
	}
	return nums
}

// @lc code=end
