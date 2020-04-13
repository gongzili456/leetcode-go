// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 34. 在排序数组中查找元素的第一个和最后一个位置
package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	res := searchRange(nums, 8)
	fmt.Println("res: ", res)
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}

	if nums == nil || len(nums) == 0 {
		return ans
	}

	var searchEdge func(isLeft bool) int
	searchEdge = func(isLeft bool) int {
		l, r := 0, len(nums) // 搜索范围超越边界

		for l < r {
			mid := (l + r) / 2

			// 搜索左边界时候，逐个搜到最左侧
			if nums[mid] > target || (isLeft && nums[mid] == target) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}

	left := searchEdge(true)

	// 不存在目标值时候，左边界会向右越界（所有值都比目标值小）
	// 或者所有值都比目标值大，此时 left 为零
	if left == len(nums) || nums[left] != target {
		return ans
	}
	ans[0] = left

	// 由于 r = mid，所以要 -1
	right := searchEdge(false) - 1
	ans[1] = right

	return ans
}
