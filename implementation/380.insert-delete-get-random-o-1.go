package main

import "math/rand"

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
	records map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		records: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.records[val]
	if !exists {
		this.records[val] = true
	}
	return !exists
}

func (this *RandomizedSet) Remove(val int) bool {
	_, exists := this.records[val]
	if exists {
		delete(this.records, val)
	}
	return exists
}

func (this *RandomizedSet) GetRandom() int {
	l := len(this.records)
	randNumber := rand.Intn(l)
	counter := 0

	for k := range this.records {
		if counter == randNumber {
			return k
		}
		counter++
	}
	return -1
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
