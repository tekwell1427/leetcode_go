package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
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
			preNums = append(preNums, nums[i])
			helper(i, preNums)
			preNums = preNums[:len(preNums)-1]
		}
	}
	helper(-1, []int{})
	return output
}

// @lc code=end
