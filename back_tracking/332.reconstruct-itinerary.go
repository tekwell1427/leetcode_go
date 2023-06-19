package main

import "sort"

/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

// @lc code=start
type Custype struct {
	Visited bool
	Target  string
}
type CusList []*Custype

func (a CusList) Len() int           { return len(a) }
func (a CusList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CusList) Less(i, j int) bool { return a[i].Target < a[j].Target }

func findItinerary(tickets [][]string) []string {
	record := make(map[string]CusList)
	for _, ticket := range tickets {
		depa, dest := ticket[0], ticket[1]
		record[depa] = append(record[depa], &Custype{Target: dest})
	}
	for key, value := range record {
		// sort.Slice(value, func(i, j int) bool {
		// 	return value[i][0] > value[j][0]
		// })
		sort.Sort(value)
		record[key] = value
	}
	output := []string{}
	depa := "JFK"
	output = append(output, depa)

	var helper func(depa string, itineraryCount int) bool
	helper = func(depa string, itineraryCount int) bool {
		if itineraryCount-1 == len(tickets) {
			return true
		}

		for i := 0; i < len(record[depa]); i++ {
			dest := record[depa][i]
			if dest.Visited {
				continue
			}
			dest.Visited = true
			output = append(output, dest.Target)
			itineraryCount++
			if helper(dest.Target, itineraryCount) {
				return true
			}
			dest.Visited = false
			output = output[:len(output)-1]
			itineraryCount--
		}
		return false
	}
	helper(depa, 1)

	return output
}

// @lc code=end
