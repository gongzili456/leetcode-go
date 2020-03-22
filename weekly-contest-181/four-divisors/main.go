// https://leetcode-cn.com/contest/weekly-contest-181/problems/four-divisors/
// 5178. 四因数
package main

import "fmt"

func main() {
	nums := []int{21, 4, 7}

	res := sumFourDivisors(nums)
	fmt.Println("res: ", res)
}

func sumFourDivisors(nums []int) int {
	sum := 0

	for _, n := range nums {
		count := 0
		ans := 0

		// i * i <= n
		for i := 1; i*i <= n; i++ {
			// 能被整除，即为因数
			if n%i == 0 {
				// 结果加上 i
				ans += i

				if i*i == n {
					// 计数只能 +1, 因为两数相同
					count++
				} else {
					// 计数 +2
					count += 2
					ans += n / i
				}
			}
		}

		if count == 4 {
			sum += ans
		}
	}
	return sum
}
