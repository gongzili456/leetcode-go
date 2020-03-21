// https://leetcode-cn.com/problems/triangle/
// 三角形最小路径和
package main

import "fmt"

func main() {
	nums := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	res := minimumTotal(nums)
	fmt.Println("res: ", res)
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < i; j++ {
			one := triangle[i-1][j] + triangle[i][j]
			if j-1 >= 0 {
				one = min(one, triangle[i-1][j-1]+triangle[i][j])
			}
			triangle[i][j] = one
		}
		triangle[i][i] += triangle[i-1][i-1]
	}

	candidates := triangle[len(triangle)-1]
	minimum := candidates[0]

	for _, v := range candidates {
		if minimum > v {
			minimum = v
		}
	}

	return minimum
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
