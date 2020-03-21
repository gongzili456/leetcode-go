// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 数组中的第K个最大元素
package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4, 6}
	res := findKthLargest(nums, 5)
	fmt.Println("res: ", res)
}

func findKthLargest(nums []int, k int) int {
	h := &maxNum{}
	for _, n := range nums {
		heap.Push(h, n)
	}

	fmt.Println("h: ", h)

	kth := 0
	for k > 0 {
		kth = heap.Pop(h).(int)
		fmt.Println("KTH: ", kth)
		k--
	}
	return kth
}

type maxNum []int

func (m maxNum) Len() int           { return len(m) }
func (m maxNum) Less(i, j int) bool { return m[i] > m[j] }
func (m maxNum) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *maxNum) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *maxNum) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
