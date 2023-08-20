package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	output := [][]int{}

	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// prune
		if nums[first] > 0 {
			break
		}
		third := len(nums) - 1
		second := first + 1
		for second < third {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if third < len(nums)-1 && nums[third] == nums[third+1] {
				third--
				continue
			}
			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				output = append(output, []int{nums[first], nums[second], nums[third]})
				second++
				third--
			} else if sum > 0 {
				third--
			} else {
				second++
			}
		}
	}
	return output
}

// @lc code=end
