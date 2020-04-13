// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// 面试题10- II. 青蛙跳台阶问题
package main

import "fmt"

func main() {
	res := numWays(7)
	fmt.Println("res: ", res)
}

func numWays(n int) int {
	if n == 1 {
		return 1
	}

	a, b := 1, 1

	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}

	return b
}
