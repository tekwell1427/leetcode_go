package main

/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Number of Provinces
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	parents := make([]int, len(isConnected))
	rank := make([]int, len(isConnected))

	for i := range parents {
		parents[i] = i
		rank[i] = 1
	}

	var find func(nodeIdx int) int
	find = func(nodeIdx int) int {
		if parents[nodeIdx] != nodeIdx {
			parents[nodeIdx] = find(parents[nodeIdx])
		}
		return parents[nodeIdx]
	}
	union := func(node1Idx, node2Idx int) bool {
		isCircle := false
		parent1 := find(node1Idx)
		parent2 := find(node2Idx)
		if parent1 == parent2 {
			isCircle = true
		} else {
			if rank[parent1] > rank[parent2] {
				parents[parent2] = parent1
			}
			if rank[parent1] < rank[parent2] {
				parents[parent1] = parent2
			}
			if rank[parent1] == rank[parent2] {
				parents[parent1] = parent2
				rank[parent2]++
			}
		}
		return isCircle
	}
	for curIdx, curConnections := range isConnected {
		// prune connectIdx is from curIdx +1, not from 0
		for connectIdx := curIdx + 1; connectIdx < len(isConnected); connectIdx++ {
			if curConnections[connectIdx] == 1 {
				union(curIdx, connectIdx)
			}
		}
	}
	numProvince := 0
	for parentIdx := range parents {
		if parents[parentIdx] == parentIdx {
			numProvince++
		}
	}
	return numProvince
}

// @lc code=end
