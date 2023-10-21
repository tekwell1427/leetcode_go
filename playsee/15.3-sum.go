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

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		one := nums[i]
		j, z := i+1, len(nums)-1

		for j < z {
			if j > i+1 && nums[j-1] == nums[j] {
				j++
				continue
			}
			two, three := nums[j], nums[z]
			sum := one + two + three
			if sum == 0 {
				output = append(output, []int{one, two, three})
				j++
			} else if sum > 0 {
				for {
					z--
					if z < 0 || nums[z] != nums[z+1] {
						break
					}
				}
			} else {
				j++
			}
		}
	}
	return output
}

// @lc code=end
