// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
// 面试题04. 二维数组中的查找
package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	res := findNumberIn2DArray(matrix, 20)

	fmt.Println("res: ", res)
}

// 暴力法：直接遍历所有元素
// 线性扫描：从左上角开始
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	h, w := len(matrix), len(matrix[0])

	x, y := 0, w-1

	for x < h && y >= 0 {
		c := matrix[x][y]

		if c == target {
			return true
		} else if c > target {
			y--
		} else {
			x++
		}
	}

	return false
}
