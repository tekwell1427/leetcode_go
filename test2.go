package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}

	// 使用 reduce 对列表元素进行求和操作
	sum := reduce(data, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println("Sum:", sum) // 输出：Sum: 15

	// 使用 map 对列表元素进行平方操作
	squared := myMap(data, func(x int) int {
		return x * x
	})
	fmt.Println("Squared:", squared) // 输出：Squared: [1 4 9 16 25]
}

func reduce(data []int, fn func(int, int) int, initial int) int {
	result := initial
	for _, value := range data {
		result = fn(result, value)
	}
	return result
}

func myMap(data []int, fn func(int) int) []int {
	result := make([]int, len(data))
	for i, value := range data {
		result[i] = fn(value)
	}
	return result
}
