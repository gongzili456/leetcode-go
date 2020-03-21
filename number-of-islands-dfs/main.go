// https://leetcode-cn.com/problems/number-of-islands/
// 200. 岛屿数量
package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	res := numIslands(grid)
	fmt.Println("res: ", res)
}

// 利用深度优先搜索 DFS，
func numIslands(grid [][]byte) int {
	H := len(grid)
	if H == 0 {
		return 0
	}

	W := len(grid[0])

	dirs := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= H || j < 0 || j >= W || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'

		// 向上下左右四个方向搜
		for _, d := range dirs {
			i1, j1 := d[0]+i, d[1]+j
			dfs(i1, j1)
		}
	}

	count := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}
