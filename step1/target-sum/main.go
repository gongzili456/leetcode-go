// https://leetcode-cn.com/problems/target-sum/
// 494. 目标和
package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}

	res := findTargetSumWays(nums, 3)
	fmt.Println("res: ", res)
}

func findTargetSumWays(nums []int, S int) int {
	count := 0
	var dfs func(sum, depth int)
	dfs = func(sum, depth int) {
		if depth == len(nums) {
			if sum == S {
				count++
			}
			return
		}

		dfs(sum-nums[depth], depth+1)
		dfs(sum+nums[depth], depth+1)
	}

	dfs(nums[0], 1)
	dfs(-nums[0], 1)

	return count
}
