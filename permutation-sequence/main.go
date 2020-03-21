// 第 K 个排列
package main

import "fmt"

func main() {
	res := getPermutation(4, 9)
	fmt.Println("res: ", res)
}

// 将 n! 种排列分为：n 组，每组有 (n - 1)!个排列，
// 根据k值可以确定是第几组的第几个排列，
// 选取该排列的第1个数字，然后递归从剩余的数字里面选取下一个数字，直到n=1为止。
func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	flag := make([]bool, n)

	var dfs func(n, k int) string

	dfs = func(n, k int) string {
		if n == 0 {
			return ""
		}

		f := factorial(n - 1)
		c := k / f

		fmt.Println("k: ", k, "f: ", f, "c: ", c)

		index := 0
		for i, j := 0, c; j >= 0; i++ {
			if flag[i] == false {
				j--
				index = i
			}
		}
		flag[index] = true
		res := string(byte(index+1) + '0')
		res += dfs(n-1, k-c*f)
		return res
	}

	return dfs(n, k-1)
}

func factorial(i int) int {
	res := 1

	for i > 1 {
		res *= i
		i--
	}

	return res
}
