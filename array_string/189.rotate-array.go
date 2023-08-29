package main

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	revert(nums)
	k = k % len(nums)
	revert(nums[:k])
	revert(nums[k:])

}
func revert(a []int) {
	left, right := 0, len(a)-1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}

// @lc code=end
