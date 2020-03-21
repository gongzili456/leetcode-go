// 最长连续序列
package main

import "fmt"

func main() {
	nums := []int{
		100, 4, 200, 1, 3, 2,
	}

	result := longestConsecutive(nums)
	fmt.Println("result: ", result)
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n] = n
	}

	ans := 0
	for _, n := range nums {
		var ok bool
		_, ok = m[n-1]

		if ok {
			continue
		}

		fmt.Println("curr: ", n, n-1, ok)

		max := 1
		n++
		_, ok = m[n]
		for ok {
			max++
			n++
			_, ok = m[n]
		}
		if ans < max {
			ans = max
		}
	}

	return ans
}
