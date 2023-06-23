package main

import "fmt"

func MergeSort(s []int) []int {
	var helper func(left, right int) []int
	helper = func(left, right int) []int {
		if left == right {
			return []int{s[left]}
		}
		mid := (left + right) / 2
		l := helper(left, mid)
		r := helper(mid+1, right)

		output := []int{}
		i, j := 0, 0

		for i < len(l) && j < len(r) {
			if l[i] > r[j] {
				output = append(output, r[j])
				j++
			} else {
				output = append(output, l[i])
				i++
			}
		}
		for ; i < len(l); i++ {
			output = append(output, l[i])
		}
		for ; j < len(r); j++ {
			output = append(output, r[j])
		}
		return output
	}
	return helper(0, len(s)-1)
}

func main() {
	a := []int{1, 5, 7, 8, 4, 9, 10}
	a = MergeSort(a)
	fmt.Println(a)
}
