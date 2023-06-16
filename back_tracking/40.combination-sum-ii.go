package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	output := [][]int{}
	var helper func(preIdx int, preNums []int, preSum int)
	helper = func(preIdx int, preNums []int, preSum int) {
		// preCandidate := -1
		for i := preIdx + 1; i < len(candidates); i++ {
			if candidates[i]+preSum > target {
				break
			}
			if i > preIdx+1 && candidates[i] == candidates[i-1] {
				continue
			}
			preNums = append(preNums, candidates[i])
			preSum += candidates[i]
			fmt.Println(preNums, preSum)

			if preSum == target {

				tmp := make([]int, len(preNums))
				copy(tmp, preNums)
				output = append(output, tmp)
			} else if preSum < target {
				helper(i, preNums, preSum)
			}
			preNums = preNums[:len(preNums)-1]
			preSum -= candidates[i]
		}
	}
	helper(-1, []int{}, 0)
	return output
}

// @lc code=end
