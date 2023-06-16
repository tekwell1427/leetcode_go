package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	fmt.Println(candidates)
	output := [][]int{}
	var helper func(preIdx int, preNums []int, preSum int)
	helper = func(preIdx int, preNums []int, preSum int) {
		// preCandidate := -1
		for i := preIdx; i < len(candidates); i++ {
			if candidates[i]+preSum > target {
				break
			}
			if i > 0 && candidates[i] == candidates[i-1] {
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
	helper(0, []int{}, 0)
	return output
}

// @lc code=end
