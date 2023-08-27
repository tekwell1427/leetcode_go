package main

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		tail1, tail2 := nums1[m-1], nums2[n-1]
		if tail1 > tail2 {
			nums1[n+m-1] = tail1
			m--
		} else {
			nums1[n+m-1] = tail2
			n--
		}
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}

// @lc code=end
