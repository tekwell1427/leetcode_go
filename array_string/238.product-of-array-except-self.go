package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	leftProduct := make([]int, len(nums))
	leftProduct[0] = 1

	for numIdx := 1; numIdx < len(nums); numIdx++ {
		leftProduct[numIdx] = leftProduct[numIdx-1] * nums[numIdx-1]
	}
	rightProduct := 1
	for numIdx := len(nums) - 2; numIdx >= 0; numIdx-- {
		rightProduct = rightProduct * nums[numIdx+1]
		leftProduct[numIdx] *= rightProduct
	}
	return leftProduct
}

// @lc code=end
