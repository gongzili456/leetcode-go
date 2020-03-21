// 5356. 矩阵中的幸运数
package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12},
	}
	res := luckyNumbers(matrix)
	fmt.Println("res: ", res)
}

// 找出最小，和最大两个数组，同时出现在两个数组的数字即满足同行最小和同列最大
func luckyNumbers(matrix [][]int) []int {
	width, height := len(matrix[0]), len(matrix)

	mr := make([]int, height)
	mc := make([]int, width)

	for i := 0; i < height; i++ {
		mr[i] = int(^uint(0) >> 1)
		for j := 0; j < width; j++ {
			if mr[i] > matrix[i][j] {
				mr[i] = matrix[i][j]
			}
			if mc[j] < matrix[i][j] {
				mc[j] = matrix[i][j]
			}
		}
	}

	ans := []int{}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if mr[i] == mc[j] {
				ans = append(ans, mr[i])
			}
		}
	}

	return ans
}
