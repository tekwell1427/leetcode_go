package main

import "fmt"

func main() {
	stoneWeights := []int{1, 3, 4}
	stoneValues := []int{15, 20, 30}
	bagWeight := 4

	dp := make([]int, bagWeight+1)

	for stoneIdx, stoneWeight := range stoneWeights {
		for bagCurWeight := bagWeight; bagCurWeight >= 0; bagCurWeight-- {
			if bagCurWeight >= stoneWeight {
				dp[bagCurWeight] = max(dp[bagCurWeight], dp[bagCurWeight-stoneWeight]+stoneValues[stoneIdx])
			}
		}
	}
	fmt.Println(dp)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
