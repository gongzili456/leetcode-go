// https://leetcode-cn.com/problems/maximal-square/solution/li-jie-san-zhe-qu-zui-xiao-1-by-lzhlyle/
// 最大正方形
package main

import "fmt"

func main() {
	matrix := [][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	res := maximalSquare(matrix)
	fmt.Println("res: ", res)
}

func maximalSquare(matrix [][]byte) int {
	if matrix == nil {
		return 0
	}

	if len(matrix) <= 0 {
		return 0
	}
	if len(matrix[0]) <= 0 {
		return 0
	}

	height := len(matrix)
	width := len(matrix[0])

	dp := make([][]byte, height+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]byte, width+1)
	}

	max := byte(0)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j], dp[i][j]) + 1
				if dp[i+1][j+1] > max {
					max = dp[i+1][j+1]
				}
			}
		}
	}

	return int(max * max)
}

func min(x, y, z byte) byte {
	if x > y {
		x = y
	}

	if x > z {
		x = z
	}

	return x
}
