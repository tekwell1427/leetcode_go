package main

/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	output := [][]int{}

	var helper func(nodeIdx int, subOutput []int)
	helper = func(nodeIdx int, subOutput []int) {
		if nodeIdx == len(graph)-1 {
			subOutput = append(subOutput, nodeIdx)
			output = append(output, subOutput)
			return
		}
		subOutput = append(subOutput, nodeIdx)

		for _, g := range graph[nodeIdx] {
			tmp := make([]int, len(subOutput))
			copy(tmp, subOutput)
			helper(g, tmp)
		}
	}
	helper(0, []int{})
	return output
}

// @lc code=end
