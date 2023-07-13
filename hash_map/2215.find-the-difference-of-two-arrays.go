package main

/*
 * @lc app=leetcode id=2215 lang=golang
 *
 * [2215] Find the Difference of Two Arrays
 */

// @lc code=start
func findDifference(nums1 []int, nums2 []int) [][]int {
	record1 := make(map[int]bool)
	record2 := make(map[int]bool)

	for _, num := range nums1 {
		record1[num] = true
	}
	for _, num := range nums2 {
		record2[num] = true
	}
	unique1, unique2 := []int{}, []int{}
	for key, _ := range record1 {
		if !record2[key] {
			unique1 = append(unique1, key)
		}
	}
	for key, _ := range record2 {
		if !record1[key] {
			unique2 = append(unique2, key)
		}
	}
	return [][]int{unique1, unique2}
}

// @lc code=end
