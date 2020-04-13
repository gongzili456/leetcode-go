// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
// 面试题12. 矩阵中的路径
package main

import "fmt"

func main() {
	board := [][]byte{
		// {'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
		// {'a', 'b'},
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'},
	}
	// word := "ABCCED"
	// word := "ba"
	word := "ABCESEEEFS"

	res := exist(board, word)
	fmt.Println("res: ", res)
}

func exist(board [][]byte, word string) bool {
	X, Y := len(board), len(board[0])

	dirs := [][]int{
		{0, -1}, // left
		{0, 1},  // right
		{-1, 0}, // up
		{1, 0},  // down
	}

	seen := map[[2]int]bool{}
	ans := false

	var search func(x, y, step int)
	search = func(x, y, step int) {
		if ans {
			return
		}

		if board[x][y] != word[step] {
			return
		}

		if step == len(word)-1 {
			ans = true
			return
		}

		pp := [2]int{x, y}
		seen[pp] = true

		for _, d := range dirs {
			x1, y1 := x+d[0], y+d[1]

			if x1 < 0 || x1 >= X || y1 < 0 || y1 >= Y {
				continue
			}
			p := [2]int{x1, y1}
			if _, ok := seen[p]; !ok {
				search(x1, y1, step+1)
			}
		}
		delete(seen, pp)
	}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if ans {
				break
			}
			search(i, j, 0)
		}
	}

	return ans
}
