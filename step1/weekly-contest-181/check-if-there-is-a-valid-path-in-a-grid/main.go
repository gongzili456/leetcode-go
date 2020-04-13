// https://leetcode-cn.com/contest/weekly-contest-181/problems/check-if-there-is-a-valid-path-in-a-grid/
// 5366. 检查网格中是否存在有效路径
package main

import "fmt"

func main() {
	grid := [][]int{
		// {2, 4, 3}, {6, 5, 2},
		{1, 2, 1}, {1, 2, 1},
	}

	res := hasValidPath(grid)
	fmt.Println("res: ", res)
}

func hasValidPath(grid [][]int) bool {
	dirs := map[int][][]int{
		1: {
			// x, y
			{0, 1}, {0, -1}, // 1, 左右
		},
		// 2, up down
		2: {
			{-1, 0}, {1, 0},
		},
		// left, down
		3: {
			{0, -1}, {1, 0},
		},
		// 4, right, down
		4: {
			{0, 1}, {1, 0},
		},
		// left, up
		5: {
			{0, -1}, {-1, 0},
		},
		// up, right
		6: {
			{-1, 0}, {0, 1},
		},
	}

	stack := []int{}

	stack = append(stack, 0, 0)
	H, W := len(grid), len(grid[0])
	seen := map[[2]int]bool{}

	for len(stack) > 0 {
		// pop
		n := len(stack)
		// 注意 x,y 的出站顺序
		y, x := stack[n-1], stack[n-2]
		stack = stack[:n-2]

		if x == H-1 && y == W-1 {
			return true
		}

		pp := [2]int{x, y}
		seen[pp] = true

		// 搜索方向
		ds := dirs[grid[x][y]]
		for _, d := range ds {
			x1, y1 := x+d[0], y+d[1]

			// 检查越界
			if x1 < 0 || x1 >= H || y1 < 0 || y >= W {
				continue
			}

			// 检查是否已经走过
			p := [2]int{x1, y1}
			if _, ok := seen[p]; ok {
				continue
			}

			// 检查是否可以返回
			nds := dirs[grid[x1][y1]]
			for _, nd := range nds {
				if d[0]+nd[0] == 0 && d[1]+nd[1] == 0 {
					stack = append(stack, x1, y1)
				}
			}
		}
		fmt.Println("s: ", stack)
	}

	return false
}
