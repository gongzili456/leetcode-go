// 5359. 最大的团队表现值
package main

import (
	"container/heap"
	"fmt"
)

func main() {
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	res := maxPerformance(6, speed, efficiency, 3)
	fmt.Println("res: ", res)
}

// 利用最小堆计算和，最大堆倒序寻找最小的乘数
func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	eh := &minNum{}

	for i, e := range efficiency {
		e <<= 32
		heap.Push(eh, -(e + i))
	}

	sh := &minNum{}
	ans := 0
	sum := 0
	added := 0

	for eh.Len() > 0 {
		ev := -heap.Pop(eh).(int)
		i := ev % (1 << 31)
		e := ev >> 32

		heap.Push(sh, speed[i])
		sum += speed[i]
		added++

		if added > k {
			s := heap.Pop(sh).(int)
			sum -= s
		}

		mut := e * sum
		if mut > ans {
			ans = mut
		}
	}
	return ans % 1000000007
}

type minNum []int

func (m minNum) Len() int           { return len(m) }
func (m minNum) Less(i, j int) bool { return m[i] < m[j] }
func (m minNum) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minNum) Push(x interface{}) { *m = append(*m, x.(int)) }
func (m *minNum) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
