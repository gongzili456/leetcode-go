// 5354. 通知所有员工所需的时间
package main

import "fmt"

func main() {
	res := numOfMinutes(11, 4, []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}, []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337})
	fmt.Println("res: ", res)
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	tree := map[int][]int{}

	for i, t := range manager {
		tree[t] = append(tree[t], i)
	}

	fmt.Println("tree: ", tree)

	var search func(h int) int

	search = func(h int) int {
		max := 0
		for _, v := range tree[h] {
			t := search(v)
			if max < t {
				max = t
			}
		}
		return max + informTime[h]
	}

	return search(headID)
}
