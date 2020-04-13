// 5383. 给 N x 3 网格图涂色的方案数
package main

import "fmt"

func main() {
	n := 4999
	res := numOfWays(n)
	fmt.Println("> ", res)
}

func numOfWays(n int) int {
	const mod = 1e9 + 7

	memo := map[[4]int]int{}
	var search func(i int, c1, c2, c3 int) int
	search = func(i int, c1, c2, c3 int) int {
		if i == 0 {
			return 1
		}
		k := [4]int{i, c1, c2, c3}
		if v, ok := memo[k]; ok {
			return v
		}

		count := 0

		for n1 := 1; n1 <= 3; n1++ {
			if n1 == c1 {
				continue
			}
			for n2 := 1; n2 <= 3; n2++ {
				if n2 == n1 || n2 == c2 {
					continue
				}
				for n3 := 1; n3 <= 3; n3++ {
					if n3 == c3 || n3 == n2 {
						continue
					}
					count = (count + search(i-1, n1, n2, n3)) % mod
				}
			}
		}

		memo[k] = count
		return count
	}

	return search(n, 0, 0, 0)
}
