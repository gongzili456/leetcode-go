package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 2
	s1 := "aa"
	s2 := "da"
	evil := "b"

	res := findGoodStrings(n, s1, s2, evil)
	fmt.Println("res: ", res)
}

// LeetCode 用户 tsholmes 的答案
// https://leetcode.com/tsholmes/
func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	pres := make([]int, len(evil))
	for i := 2; i < len(evil); i++ {
		for j := 1; j < i; j++ {
			if strings.HasPrefix(evil, evil[j:i]) {
				pres[i] = i - j
				break
			}
		}
	}
	fmt.Println("pres: ", pres)

	const mod = 1e9 + 7

	type state struct {
		i, m   int
		gt, lt bool
	}

	memo := map[state]int{}
	var search func(i int, m int, gt bool, lt bool) int
	search = func(i int, m int, gt bool, lt bool) int {
		if m == len(evil) {
			return 0
		}
		if i == n {
			return 1
		}
		k := state{i, m, gt, lt}
		if v, ok := memo[k]; ok {
			return v
		}

		res := 0

		min, max := s1[i], s2[i]
		if gt {
			min = 'a'
		}
		if lt {
			max = 'z'
		}

		for c := min; c <= max; c++ {
			m2 := m
			for m2 > 0 && evil[m2] != c {
				m2 = pres[m2]
			}
			if c == evil[m2] {
				m2++
			}
			res = (res + search(i+1, m2, gt || c > min, lt || c < max)) % mod
		}
		memo[k] = res
		return res
	}

	return search(0, 0, false, false)
}
