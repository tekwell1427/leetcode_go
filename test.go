package main

import (
	"fmt"
)

func main() {
	// a, _ := strconv.Atoi(string('9'))
	// fmt.Println(a, string('t'))
	// fmt.Println('a')
	// n := []int{1, 2, 3}
	// a := &myHeap{nums: n}
	// for i := 4; i < 10; i++ {
	// 	a.Push(i)
	// 	fmt.Println("push", i)
	// 	fmt.Println(a.nums)
	// 	// fmt.Println("pop", i)
	// 	// a.Pop()
	// 	// fmt.Println(a.nums)
	// }
	var m myHeap
	var t *int
	m.Push(1)
	fmt.Println(&t)
	fmt.Printf("%p\n", t)
	fmt.Printf("%p, %p  \n", &m, m.nums)
	m.test()
	// fmt.Println(m.nums)

}

type myHeap struct {
	nums []int
}

func (m *myHeap) Push(x any) {
	// fmt.Printf("2 %p\n", m.nums)
	m.nums = append(m.nums, x.(int))
}
func (m *myHeap) Pop() any {
	p := m.nums[len(m.nums)-1]
	m.nums = m.nums[:len(m.nums)-1]
	return p
}
func (m myHeap) test() {
	fmt.Printf("%p, %p", &m, m.nums)
}
