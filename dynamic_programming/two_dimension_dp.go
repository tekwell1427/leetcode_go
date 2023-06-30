package main

func main() {

	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagMaxWeight := 4
	dp := make([][]int, len(weight))
	for row := 0; row < len(dp); row++ {
		dp[row] = make([]int, bagMaxWeight+1)
		// dp[row][0] = 0
	}
	for curBagWeight := 0; curBagWeight <= bagMaxWeight; curBagWeight++ {
		if curBagWeight >= weight[0] {
			dp[0][curBagWeight] = value[0]
		}
	}
	for objectIdx := 1; objectIdx < len(weight); objectIdx++ {
		for curBagWeight := 0; curBagWeight <= bagMaxWeight; curBagWeight++ {
			if curBagWeight < weight[objectIdx] {
				dp[objectIdx][curBagWeight] = dp[objectIdx-1][curBagWeight]
			} else {
				dp[objectIdx][curBagWeight] = max(dp[objectIdx-1][curBagWeight], dp[objectIdx-1][curBagWeight-weight[objectIdx]]+value[objectIdx])
			}
		}
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
