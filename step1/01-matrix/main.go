// https://leetcode-cn.com/problems/01-matrix/
// 542. 01 矩阵
package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	res := updateMatrix(matrix)

	fmt.Println("res: ", res)
}

// 从为 0 的节点开始广度搜索 BFS，可以避免出现最差情况（每次都要遍历整个 matrix）
func updateMatrix(matrix [][]int) [][]int {
	Q := []int{}

	H, W := len(matrix), len(matrix[0])

	maxValue := int(^uint(0) >> 1)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if matrix[i][j] != 0 {
				// 将非0的节点置为 超大数
				matrix[i][j] = maxValue
			} else {
				Q = append(Q, i, j)
			}
		}
	}

	dirs := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	for len(Q) > 0 {
		r, c := Q[0], Q[1]
		Q = Q[2:]

		for _, d := range dirs {
			r1, c1 := r+d[0], c+d[1]
			if r1 < 0 || r1 >= H || c1 < 0 || c1 >= W {
				continue
			}

			// 最小距离是1，如果新方向上的值比当前节点的值 大 1，则说明碰到了非0边界
			if matrix[r1][c1] >= matrix[r][c]+1 {
				matrix[r1][c1] = matrix[r][c] + 1
				Q = append(Q, r1, c1)
			}
		}
	}

	return matrix
}
