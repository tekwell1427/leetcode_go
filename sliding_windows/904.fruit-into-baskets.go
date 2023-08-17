package main

/*
 * @lc app=leetcode id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 */

// @lc code=start
func totalFruit(fruits []int) int {
	record := make(map[int]int)
	left := 0
	output := 0
	for right, rightFruit := range fruits {
		record[rightFruit]++
		for len(record) > 2 {
			leftFruit := fruits[left]
			record[leftFruit]--
			if record[leftFruit] == 0 {
				delete(record, leftFruit)
			}
			left++
		}
		output = max(output, right-left+1)
	}
	return output
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
