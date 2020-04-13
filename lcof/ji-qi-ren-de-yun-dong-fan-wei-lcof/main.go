// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
// 面试题13. 机器人的运动范围
package main

import "fmt"

func main() {
	m, n, k := 2, 3, 1
	res := movingCount(m, n, k)
	fmt.Println("res: ", res)
}

func movingCount(m int, n int, k int) int {
	dirs := [][]int{
		{0, -1}, // left
		{0, 1},  // right
		{-1, 0}, // up
		{1, 0},  // down
	}

	ans := 0
	seen := map[[2]int]bool{}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		s := sum(x, y)
		if s > k {
			return
		}

		ans++

		pp := [2]int{x, y}
		seen[pp] = true

		for _, d := range dirs {
			x1, y1 := x+d[0], y+d[1]
			if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
				continue
			}
			p := [2]int{x1, y1}
			if _, ok := seen[p]; !ok {
				dfs(x1, y1)
			}
		}
	}

	dfs(0, 0)
	return ans
}

func sum(x, y int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	for y > 0 {
		sum += y % 10
		y /= 10
	}
	return sum
}
