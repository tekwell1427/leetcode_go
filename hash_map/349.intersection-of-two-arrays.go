package main

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	output := []int{}
	record := make(map[int]bool)

	for _, n := range nums1 {
		record[n] = true
	}
	for _, n := range nums2 {
		if _, exits := record[n]; exits {
			output = append(output, n)
			delete(record, n)
		}
	}
	return output
}

// @lc code=end
