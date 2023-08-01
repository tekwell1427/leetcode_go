package main

import "fmt"

/*
 * @lc app=leetcode id=684 lang=golang
 *
 * [684] Redundant Connection
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	// +1: because node number is from 1 to edges.length
	parents := make([]int, len(edges)+1)
	rank := make([]int, len(edges)+1)
	// initially, node point to itself
	for i := range parents {
		parents[i] = i
		rank[i] = 1
	}
	fmt.Println(parents, rank)
	// find node parentIdx
	var find func(nodeIdx int) int
	find = func(nodeIdx int) int {
		if parents[nodeIdx] != nodeIdx {
			nodeIdx = find(parents[nodeIdx])
		}
		return parents[nodeIdx]
	}
	union := func(node1Idx, node2Idx int) bool {
		var hasCircle bool
		node1Parent := find(node1Idx)
		node2Parent := find(node2Idx)
		if node1Parent == node2Parent {
			hasCircle = true
		} else {
			if rank[node1Parent] > rank[node2Parent] {
				parents[node2Parent] = node1Parent
			} else if rank[node1Parent] < rank[node2Parent] {
				parents[node1Parent] = node2Parent
			} else {
				parents[node1Parent] = node2Parent
				rank[node2Parent]++
			}
			hasCircle = false
		}
		// fmt.Println(node1Parent, node2Parent, parents, node1Idx, node2Idx)
		return hasCircle
	}
	for _, edge := range edges {
		if union(edge[0], edge[1]) {
			return edge
		}
	}
	return []int{}
}

// @lc code=end
