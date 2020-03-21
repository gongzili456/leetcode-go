// 朋友圈
package main

import "fmt"

func main() {
	nums := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 1, 1},
	}

	res := findCircleNum(nums)
	fmt.Println("res: ", res)
}

// 无向图，邻接矩阵
// 一维数组表示顶点，二维数组表示边
func findCircleNum(M [][]int) int {
	// M 为边
	visited := make([]int, len(M)) // 顶点

	count := 0
	// 深度优先搜索
	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			dfs(M, visited, i) // 最浅层
			count++
		}
	}
	return count
}

// i 是当前顶点
func dfs(M [][]int, visited []int, i int) {
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 && visited[j] == 0 {
			visited[j] = 1
			dfs(M, visited, j)
		}
	}
}
