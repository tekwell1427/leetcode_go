package main

/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	record := make(map[int]int, len(nums2))

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			preNum := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			record[preNum] = num
		}
		stack = append(stack, num)
	}
	output := make([]int, len(nums1))
	for i, num := range nums1 {
		if key, exits := record[num]; exits {
			output[i] = key
		} else {
			output[i] = -1
		}
	}
	return output
}

// @lc code=end
