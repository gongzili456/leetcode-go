// https://leetcode-cn.com/problems/longest-increasing-subsequence/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=transfer2china
// 300. 最长上升子序列
package main

import "fmt"

func main() {
	nums := []int{
		10, 9, 2, 5, 3, 7, 101, 18,
	}
	res := lengthOfLIS(nums)
	fmt.Println("res: ", res)
}

// 动态规划方法:DP，比较 nums[i] 与之前所有的数字，
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 1
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		maxVal := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && maxVal < dp[j] {
				maxVal = dp[j]
			}
		}
		dp[i] = maxVal + 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
