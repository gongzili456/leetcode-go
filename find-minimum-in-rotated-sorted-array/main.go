// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
// 153. 寻找旋转排序数组中的最小值
package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	res := findMin(nums)
	fmt.Println("res: ", res)
}

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1

	for l <= r {
		if l == r {
			return nums[l]
		}

		mid := (l + r) / 2

		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
