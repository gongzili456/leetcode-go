// https://leetcode-cn.com/problems/number-of-islands/
// 200. 岛屿数量
package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}

	res := numIslands(grid)
	fmt.Println("res: ", res)
}

// 利用广度优先搜索 BFS，
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

	var bfs func(i, j int)
	bfs = func(i, j int) {
		q := []int{}
		q = append(q, i, j)

		for len(q) > 0 {
			i, j := q[0], q[1]
			q = q[2:]
			// 向上下左右四个方向搜
			for _, d := range dirs {
				i1, j1 := d[0]+i, d[1]+j
				if i1 >= 0 && i1 < H && j1 >= 0 && j1 < W && grid[i1][j1] == '1' {
					grid[i1][j1] = '0'
					q = append(q, i1, j1)
				}
			}
		}
	}

	count := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == '1' {
				count++
				bfs(i, j)
			}
		}
	}

	return count
}
