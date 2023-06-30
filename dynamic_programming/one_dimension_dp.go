package main

import "fmt"

func main() {

	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagMaxWeight := 4
	dp := make([]int, bagMaxWeight+1)

	for objectIdx := 0; objectIdx < len(weight); objectIdx++ {
		for curBagWeight := bagMaxWeight; curBagWeight >= weight[objectIdx]; curBagWeight-- {
			dp[curBagWeight] = max(dp[curBagWeight], dp[curBagWeight-weight[objectIdx]]+value[objectIdx])
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
