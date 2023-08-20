package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for idx, num := range nums {
		if value, exits := record[target-num]; exits {
			return []int{value, idx}
		}
		record[num] = idx
	}
	return []int{}
}

// @lc code=end
