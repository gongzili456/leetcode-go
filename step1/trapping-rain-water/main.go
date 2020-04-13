// 接雨水
// https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/
package main

import "fmt"

func main() {
	n := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(n)
	fmt.Println("rs: ", res)
}

func trap(height []int) int {
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1

	ans := 0

	for left < right {
		if height[left] < height[right] {
			if leftMax > height[left] {
				ans += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if rightMax > height[right] {
				ans += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return ans
}
