package main

import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	output := [][]int{}
	var helper func(preIdx int, preNums []int)
	helper = func(preIdx int, preNums []int) {
		// append
		tmp := make([]int, len(preNums))
		copy(tmp, preNums)
		output = append(output, tmp)

		if preIdx == len(nums)-1 {
			return
		}

		for i := preIdx + 1; i < len(nums); i++ {
			if i > preIdx+1 && nums[i] == nums[i-1] {
				continue
			}
			preNums = append(preNums, nums[i])
			helper(i, preNums)
			preNums = preNums[:len(preNums)-1]
		}
	}
	helper(-1, []int{})
	return output
}

// @lc code=end
