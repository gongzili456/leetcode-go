package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 3, 3}
	res := findLucky(arr)
	fmt.Println("res: ", res)
}

func findLucky(arr []int) int {
	m := map[int]int{}

	for _, n := range arr {
		if v, ok := m[n]; ok {
			m[n] = v + 1
		} else {
			m[n] = 1
		}
	}

	ans := -1

	for k, v := range m {
		if k == v && k > ans {
			ans = k
		}
	}

	return ans
}
