package main

import "strconv"

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	output := []string{}
	if len(nums) == 0 {
		return output
	}

	head := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 != nums[i-1] {
			if head != nums[i-1] {
				output = append(output, strconv.Itoa(head)+"->"+strconv.Itoa(nums[i-1]))
			} else {
				output = append(output, strconv.Itoa(head))
			}
			head = nums[i]
		}
	}
	if head != nums[len(nums)-1] {
		output = append(output, strconv.Itoa(head)+"->"+strconv.Itoa(nums[len(nums)-1]))
	} else {
		output = append(output, strconv.Itoa(head))
	}
	return output
}

// @lc code=end
