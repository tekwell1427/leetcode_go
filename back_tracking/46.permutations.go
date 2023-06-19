package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
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
		for i := 0; i < len(nums); i++ {
			if !used[i] {
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
