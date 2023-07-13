package main

/*
 * @lc app=leetcode id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	record := make(map[int]int)

	for _, a := range arr {
		record[a]++
	}
	countRecord := make(map[int]bool)
	for _, value := range record {

		if exist, _ := countRecord[value]; exist {
			return false
		}
		countRecord[value] = true
	}
	return true
}

// @lc code=end
