package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 4, 7, 6, 7}

	res := minSubsequence(nums)
	fmt.Println("> ", res)
}

func minSubsequence(nums []int) []int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	ans := []int{}
	max := 0

	for i := 0; i < len(nums); i++ {
		if max <= sum-max {
			max += nums[i]
			ans = append(ans, nums[i])
		}
	}
	return ans
}
