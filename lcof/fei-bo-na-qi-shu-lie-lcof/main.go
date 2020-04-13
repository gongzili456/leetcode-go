// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// 面试题10- I. 斐波那契数列
package main

import "fmt"

func main() {
	res := fib(45)
	fmt.Println("res: ", res)
}

func fib(n int) int {
	p1, p2 := 0, 1
	if n == 0 {
		return p1
	}
	if n == 1 {
		return p2
	}

	// fmt.Printf("%d, %d, ", p1, p2)

	for i := 2; i <= n; i++ {
		t := (p1 + p2) % (1e9 + 7)
		// fmt.Printf("%d, ", t)
		p1 = p2
		p2 = t
	}

	return p2
}
