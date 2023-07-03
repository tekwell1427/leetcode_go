package main

/*
 * @lc app=leetcode id=1679 lang=golang
 *
 * [1679] Max Number of K-Sum Pairs
 */

// @lc code=start
func maxOperations(nums []int, k int) int {
	record := make(map[int]int)
	output := 0
	for _, num := range nums {
		if num >= k {
			continue
		}
		if v, exists := record[k-num]; exists && v > 0 {
			output++
			record[k-num]--
			continue
		}
		record[num]++
	}
	return output
}

// @lc code=end
