// https://leetcode-cn.com/problems/valid-perfect-square/
// 367. 有效的完全平方数
package main

import "fmt"

func main() {
	res := isPerfectSquare(16)
	fmt.Println("res: ", res)
}

func isPerfectSquare(num int) bool {
	l, r := 1, num

	for l < r {
		mid := (l + r) / 2
		fmt.Println("mid: ", mid)

		if mid*mid == num {
			return true
		}

		if mid*mid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l*l == num {
		return true
	}

	return false
}
