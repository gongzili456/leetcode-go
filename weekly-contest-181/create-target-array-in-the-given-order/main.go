// 5364. 按既定顺序创建目标数组
package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	res := createTargetArray(nums, index)

	fmt.Println("res: ", res)
}

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, 0)

	var add func(i, n int)
	add = func(i, n int) {
		// 后移
		t := make([]int, 0)
		af := target[i:]
		bf := target[:i]
		t = append(t, bf...)
		t = append(t, n)
		t = append(t, af...)
		target = t
	}

	for i, v := range index {
		add(v, nums[i])
	}

	return target
}
