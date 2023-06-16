package main

/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	output := [][]int{}
	var helper func(preIdx int, preSum int, preNums []int)
	helper = func(preIdx, preSum int, preNums []int) {
		if len(preNums) >= k {
			return
		}
		for i := preIdx + 1; i <= 9; i++ {
			preSum += i
			preNums = append(preNums, i)
			if preSum == n && len(preNums) == k {
				tmp := make([]int, len(preNums))
				copy(tmp, preNums)
				output = append(output, tmp)
				// pruning
			} else if preSum > n {
				break
			}
			helper(i, preSum, preNums)
			preNums = preNums[:len(preNums)-1]
			preSum -= i
		}
	}
	helper(0, 0, []int{})
	return output
}

// @lc code=end
