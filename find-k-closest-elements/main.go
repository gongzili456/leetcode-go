// https://leetcode-cn.com/problems/find-k-closest-elements/
// 658. 找到 K 个最接近的元素
package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	nums := []int{1, 1, 1, 10, 10}

	res := findClosestElements(nums, 1, 9)
	fmt.Println("Res: ", res)
}

func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[0:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	// 首先找到目标值 x
	l, r := 0, len(arr)-1
	targetIndex := 0

	for l < r {
		mid := (l + r) / 2
		if arr[mid] == x {
			targetIndex = mid
			break
		}

		if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// 可能不存在 x
	if arr[targetIndex] != x {
		targetIndex = l
	}
	fmt.Println("T: ", targetIndex)

	// 搜索 k 的范围
	// 最大搜索范围是：targetIndex - k - 1; targetIndex + k - 1
	// -1 是排除目标数字本身
	min := 0
	if targetIndex-k-1 > min {
		min = targetIndex - k - 1
	}
	max := len(arr) - 1
	if targetIndex+k-1 < max {
		max = targetIndex + k - 1
	}

	for max-min > k-1 {
		// 对比差值
		if x-arr[min] <= arr[max]-x {
			max--
		} else {
			min++
		}
	}

	return arr[min : max+1]
}
