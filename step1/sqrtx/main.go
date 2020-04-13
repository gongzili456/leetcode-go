// https://leetcode-cn.com/problems/sqrtx/
// 69. x 的平方根
package main

import "fmt"

func main() {
	res := mySqrt(9)
	fmt.Println("res: ", res)
}

func mySqrt(x int) int {
	l, r := 0, x

	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		}

		if mid*mid > x {
			r = mid - 1
		}

		if mid*mid < x {
			l = mid + 1
		}
	}
	if r < l {
		return r
	}
	return l
}
