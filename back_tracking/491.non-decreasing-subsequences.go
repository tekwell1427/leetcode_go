package main

/*
 * @lc app=leetcode id=491 lang=golang
 *
 * [491] Non-decreasing Subsequences
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	output := [][]int{}
	var helper func(preIdx int, preNums []int)
	helper = func(preIdx int, preNums []int) {

		if len(preNums) >= 2 {
			if preNums[len(preNums)-1] < preNums[len(preNums)-2] {
				return
			}
			tmp := make([]int, len(preNums))
			copy(tmp, preNums)
			output = append(output, tmp)
		}
		if preIdx == len(nums)-1 {
			return
		}
		hasCheck := make(map[int]bool)
		for i := preIdx + 1; i < len(nums); i++ {
			if i > preIdx+1 && hasCheck[nums[i]] {
				continue
			}
			hasCheck[nums[i]] = true
			preNums = append(preNums, nums[i])
			helper(i, preNums)
			preNums = preNums[:len(preNums)-1]
		}
	}
	helper(-1, []int{})
	return output
}

// @lc code=end
