// 5381. 查询带键的排列
package main

import "fmt"

func main() {
	queries := []int{3, 1, 2, 1}
	m := 5
	res := processQueries(queries, m)
	fmt.Println("> ", res)
}

func processQueries(queries []int, m int) []int {
	P := []int{}
	for i := 1; i <= m; i++ {
		P = append(P, i)
	}

	ans := []int{}

	var move func(p int)
	move = func(p int) {
		t := P[p]
		for i := p; i > 0; i-- {
			P[i] = P[i-1]
		}
		P[0] = t
	}

	var search func(p int) int
	search = func(p int) int {
		for i, n := range P {
			if p == n {
				return i
			}
		}
		return 0
	}

	for _, q := range queries {
		p := search(q)
		ans = append(ans, p)
		move(p)
	}

	return ans
}
