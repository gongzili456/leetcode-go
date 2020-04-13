// https://leetcode-cn.com/problems/water-and-jug-problem/solution/hu-dan-long-wei-liang-zhang-you-yi-si-de-tu-by-ant/
// 365. 水壶问题
package main

import "fmt"

func main() {
	x := 2
	y := 6
	z := 5

	res := canMeasureWater(x, y, z)
	fmt.Println("res: ", res)
}

func canMeasureWater(x int, y int, z int) bool {
	q := []int{}
	q = append(q, 0)

	seen := map[int]bool{}

	// BFS: 搜索方向 +x, -x, +y, -y
	for len(q) > 0 {
		// n 是当前桶里剩余的水
		n := q[0]
		q = q[1:]

		if _, ok := seen[n+x]; !ok && n+x <= x+y {
			q = append(q, n+x)
			seen[n+x] = true
		}

		if _, ok := seen[n-x]; !ok && n-x >= 0 {
			q = append(q, n-x)
			seen[n-x] = true
		}

		if _, ok := seen[n+y]; !ok && n+y <= x+y {
			q = append(q, n+y)
			seen[n+y] = true
		}

		if _, ok := seen[n-y]; !ok && n-y >= 0 {
			q = append(q, n-y)
			seen[n-y] = true
		}

		if _, ok := seen[z]; ok {
			return true
		}

		copy()
	}

	return false
}
