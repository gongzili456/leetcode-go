// 5355. T 秒后青蛙的位置
package main

import "fmt"

func main() {
	edges := [][]int{
		{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5},
	}

	res := frogPosition(7, edges, 20, 6)

	fmt.Println("res: ", res)
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// 构建无向图 map
	es := map[int][]int{}

	for _, e := range edges {
		es[e[0]] = append(es[e[0]], e[1])
		es[e[1]] = append(es[e[1]], e[0])
	}

	fmt.Println("ES: ", es)

	var fp float64
	visited := map[int]bool{}

	var search func(i int, p float64, tt int)
	search = func(i int, p float64, tt int) {
		// 如果到了次数
		if tt == t {
			if i == target {
				fp += p
			}
			// 如果仍然没有找到，概率为零
			return
		}

		visited[i] = true

		ec := 0
		for _, e := range es[i] {
			if !visited[e] {
				ec++
			}
		}

		for _, e := range es[i] {
			if !visited[e] {
				search(e, p/float64(ec), tt+1)
			}
		}

		if ec == 0 && i == target {
			fp += p
		}
		visited[i] = false
	}

	search(1, 1, 0)

	return fp
}
