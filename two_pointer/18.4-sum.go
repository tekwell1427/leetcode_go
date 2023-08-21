package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	output := [][]int{}

	for first := 0; first < len(nums)-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < len(nums)-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third := second + 1
			forth := len(nums) - 1

			for third < forth {
				if third > second+1 && nums[third] == nums[third-1] {
					third++
					continue
				}
				if forth < len(nums)-1 && nums[forth] == nums[forth+1] {
					forth--
					continue
				}
				sum := nums[first] + nums[second] + nums[third] + nums[forth]
				if sum == target {
					output = append(output, []int{nums[first], nums[second], nums[third], nums[forth]})
					third++
					forth--
				} else if sum > target {
					forth--
				} else {
					third++
				}
			}
		}
	}
	return output
}

// @lc code=end
