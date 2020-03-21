// https://leetcode-cn.com/problems/perfect-squares/solution/hua-jie-suan-fa-279-wan-quan-ping-fang-shu-by-guan/
// 279. 完全平方数
package main

import "fmt"

func main() {
	res := numSquares(285)
	fmt.Println("res: ", res)
}

// dp[i] = MIN(dp[i], dp[i - j * j] + 1)
func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		// fmt.Println("-", dp, i)

		for j := 1; i-j*j >= 0; j++ {
			r := i - j*j
			if dp[r]+1 < dp[i] {
				dp[i] = dp[r] + 1
				// fmt.Println("=", dp, j)
			}
		}
	}

	return dp[n]
}
