package main

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	output := [][]int{}
	// true: used; false: not used yet
	used := make([]bool, len(nums))
	var helper func(subNums []int)
	helper = func(subNums []int) {
		if len(subNums) == len(nums) {
			tmp := make([]int, len(subNums))
			copy(tmp, subNums)
			output = append(output, tmp)
			return
		}
		preUsedVal := -11
		for i := 0; i < len(nums); i++ {
			if i != -11 && nums[i] == preUsedVal {
				continue
			}
			if !used[i] {
				preUsedVal = nums[i]
				subNums = append(subNums, nums[i])
				used[i] = true
				helper(subNums)
				subNums = subNums[:len(subNums)-1]
				used[i] = false
			}
		}
	}
	helper([]int{})
	return output
}

// @lc code=end
