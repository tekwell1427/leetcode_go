package main

import (
	"math"
)

/*
 * @lc app=leetcode id=743 lang=golang
 *
 * [743] Network Delay Time
 */

// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {
	visited := make([]bool, n)
	paths := make([]int, n)
	g := make([][]int, n)

	for i, _ := range paths {
		paths[i] = math.MaxInt8
	}

	for i, _ := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt8
		}
	}
	// each row: source node
	// each column: target node
	for _, time := range times {
		srcNode, tarNode, time := time[0], time[1], time[2]
		g[srcNode-1][tarNode-1] = time
	}
	// have to minus 1

	curSrcNode := k - 1
	paths[curSrcNode] = 0

	for curSrcNode != -1 {
		for curTarNode := 0; curTarNode < n; curTarNode++ {
			paths[curTarNode] = min(paths[curTarNode], paths[curSrcNode]+g[curSrcNode][curTarNode])
		}
		visited[curSrcNode] = true

		curSrcNode = -1
		for nodeIdx, isVisited := range visited {
			if !isVisited {
				if curSrcNode == -1 {
					curSrcNode = nodeIdx
				}
				if paths[nodeIdx] < paths[curSrcNode] {
					curSrcNode = nodeIdx
				}
			}
		}
	}
	// line: 37 is equavalent with line:59
	// curSrcNode := k - 1
	// paths[curSrcNode] = 0

	// for i := 0; i < n; i++ {
	// 	curSrcNode := -1
	// 	for nodeIdx, isVisited := range visited {
	// 		if !isVisited {
	// 			if curSrcNode == -1 {
	// 				curSrcNode = nodeIdx
	// 			}
	// 			if paths[nodeIdx] < paths[curSrcNode] {
	// 				curSrcNode = nodeIdx
	// 			}
	// 		}
	// 	}
	// 	for curTarNode := 0; curTarNode < n; curTarNode++ {
	// 		paths[curTarNode] = min(paths[curTarNode], paths[curSrcNode]+g[curSrcNode][curTarNode])
	// 	}
	// 	visited[curSrcNode] = true
	// }

	output := -1
	for _, path := range paths {
		output = max(output, path)
	}
	if output == math.MaxInt8 {
		return -1
	}
	return output
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
