// https://leetcode-cn.com/problems/find-peak-element/solution/xun-zhao-feng-zhi-by-leetcode/
// 162. 寻找峰值
package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	res := findPeakElement(nums)
	fmt.Println("res: ", res)
}

func findPeakElement(nums []int) int {
	var search func(l, r int) int
	search = func(l, r int) int {
		if l == r {
			return l
		}

		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			return search(l, mid)
		}

		return search(mid+1, r)
	}

	return search(0, len(nums)-1)
}
