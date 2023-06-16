package main

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	output := [][]int{}
	var helper func(preIdx int, curNums []int)
	helper = func(preIdx int, curNums []int) {
		// if len(curNums) > k {
		// 	return
		// }
		for i := preIdx + 1; i <= n; i++ {
			curNums = append(curNums, i)
			if len(curNums) == k {
				tmp := make([]int, k)
				copy(tmp, curNums)
				output = append(output, tmp)
				curNums = curNums[:len(curNums)-1]
			} else if len(curNums) < k {
				helper(i, curNums)
				curNums = curNums[:len(curNums)-1]
			}
		}
	}
	helper(0, []int{})
	return output
}

// @lc code=end
