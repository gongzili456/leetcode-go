// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 面试题11. 旋转数组的最小数字
package main

import "fmt"

func main() {
	nums := []int{2, 2, 2, 0, 1}
	res := minArray(nums)
	fmt.Println("res: ", res)
}

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1

	for l < r {
		mid := (l + r) / 2

		if numbers[mid] < numbers[r] {
			r = mid
		} else if numbers[mid] == numbers[r] {
			r--
		} else {
			l = mid + 1
		}

	}

	return numbers[l]
}
