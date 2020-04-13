// https://leetcode-cn.com/problems/daily-temperatures/
// 739. 每日温度
package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}

	res := dailyTemperatures(temperatures)

	fmt.Println("res: ", res)
}

func dailyTemperatures(T []int) []int {
	stack := []int{}

	ans := make([]int, len(T))

	var pop func() (n, i int)
	pop = func() (n int, i int) {
		l := len(stack)
		n = stack[l-1]
		stack = stack[:l-1]
		i = n % (1 << 31)
		n = n >> 32
		return n, i
	}

	for i := 0; i < len(T); i++ {
		n := T[i]
		if i == 0 {
			stack = append(stack, n<<32+i)
			continue
		}

		prev, index := pop()
		for n > prev {
			fmt.Println("> ", prev, index, n, i)

			ans[index] = i - index
			if len(stack) <= 0 {
				break
			}
			prev, index = pop()
		}

		if n <= prev {
			stack = append(stack, prev<<32+index)
		}
		stack = append(stack, n<<32+i)
	}

	return ans
}
