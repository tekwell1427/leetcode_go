package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(s []int) {
	if len(s) < 2 {
		return
	}
	left, right := 0, len(s)-1
	pivot := rand.Int() % len(s)
	s[pivot], s[right] = s[right], s[pivot]

	for i := 0; i < len(s); i++ {
		if s[i] < s[right] {
			s[i], s[left] = s[left], s[i]
			left++
		}
	}
	s[left], s[right] = s[right], s[left]
	QuickSort(s[:left])
	QuickSort(s[left+1:])
}

func main() {
	a := []int{1, 5, 7, 8, 4, 9, 10}
	QuickSort(a)
	fmt.Println(a)
}
