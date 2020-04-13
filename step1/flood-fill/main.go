// https://leetcode-cn.com/problems/flood-fill/
// 733. 图像渲染
package main

import "fmt"

func main() {
	image := [][]int{
		{1, 1, 1}, {1, 1, 0}, {1, 0, 1},
		// {0, 0, 0}, {0, 1, 1},
	}

	res := floodFill(image, 1, 1, 2)
	fmt.Println("res: ", res)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	stack := []int{}

	stack = append(stack, sr, sc)

	dirs := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	V := image[sr][sc]
	H, W := len(image), len(image[0])

	visied := map[[2]int]bool{}

	for len(stack) > 0 {
		n := len(stack)
		r, c := stack[n-2], stack[n-1]
		stack = stack[:n-2]

		if _, ok := visied[[2]int{r, c}]; ok {
			continue
		}

		visied[[2]int{r, c}] = true

		image[r][c] = newColor

		for _, d := range dirs {
			r1, c1 := r+d[0], c+d[1]

			if r1 < 0 || r1 >= H || c1 < 0 || c1 >= W || image[r1][c1] != V {
				continue
			}

			stack = append(stack, r1, c1)
		}
	}

	return image
}
