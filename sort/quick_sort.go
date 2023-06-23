package main

import "fmt"

func QuickSort(s []int) {
	var helper func(left, right int)
	helper = func(left, right int) {
		if left >= right {
			return
		}
		// set pivot
		key := s[left]

		i, j := left, right
		for i < j {
			// right side move first
			for s[j] > key && i < j {
				j--
			}
			for s[i] <= key && i < j {
				i++
			}
			if i < j {
				s[i], s[j] = s[j], s[i]
			}
		}
		// i=j
		s[left] = s[i]
		s[i] = key

		helper(left, i-1)
		helper(i+1, right)
	}
	helper(0, len(s)-1)
}

func main() {
	a := []int{1, 5, 7, 8, 4, 9, 10}
	QuickSort(a)
	fmt.Println(a)
}
