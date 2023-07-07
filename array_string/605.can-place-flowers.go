package main

/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 || len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	preNum := 0

	for i := 0; i < len(flowerbed)-1; i++ {
		if preNum == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 || (i == len(flowerbed)-2 && flowerbed[i] == 0 && flowerbed[i+1] == 0) {
			n--
			preNum = 1
			if n == 0 {
				return true
			}
		} else if flowerbed[i] == 0 {
			preNum = 0
		} else {
			preNum = 1
		}
	}
	return false
}

// @lc code=end
