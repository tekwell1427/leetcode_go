package main

/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// for nuum1+nums2 sum of combination
	record := make(map[int]int)

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			record[n1+n2]++
		}
	}

	output := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if value, exits := record[-(n3 + n4)]; exits {
				output += value
			}
		}
	}
	return output
}

// @lc code=end
