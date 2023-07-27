package main

/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	output := [][]int{}
	subOutput := []int{0}

	var helper func(nodeIdx int)
	helper = func(nodeIdx int) {
		if nodeIdx == len(graph)-1 {
			tmp := make([]int, len(subOutput))
			copy(tmp, subOutput)
			output = append(output, tmp)
			return
		}

		for _, g := range graph[nodeIdx] {
			subOutput = append(subOutput, g)
			helper(g)
			subOutput = subOutput[:len(subOutput)-1]
		}
	}
	helper(0)
	return output
}

// @lc code=end
