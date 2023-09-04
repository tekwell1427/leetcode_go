package main

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int)

	for i, num := range nums {
		preIndex, exists := record[num]

		if exists {
			if i-preIndex <= k {
				return true
			}
		}
		record[num] = i
	}
	return false
}

// @lc code=end
