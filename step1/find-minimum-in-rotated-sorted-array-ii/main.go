// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
// 154. 寻找旋转排序数组中的最小值 II
package main

import "fmt"

func main() {
	// nums := []int{2, 2, 2, 0, 1}
	nums := []int{1, 3, 3}

	res := findMin(nums)
	fmt.Println("res: ", res)
}

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1

	for l <= r {
		fmt.Println("> ", l, r)
		if l == r {
			return nums[l]
		}
		mid := (l + r) / 2

		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = r - 1
		}
	}

	return -1
}
