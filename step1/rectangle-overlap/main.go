// https://leetcode-cn.com/problems/rectangle-overlap/
// 836. 矩形重叠
package main

import (
	"fmt"
)

func main() {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}

	res := isRectangleOverlap(rec1, rec2)
	fmt.Println("res: ", res)
}

// x := [0, 2] [1, 3] 重合部分为： 1-2
// y := [0, 2] [1, 3] 重合部分为：1-2
// 将数据分别映射到 X，Y 轴上，线段都重合则面积重合
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[0] >= rec2[2] || rec1[2] <= rec2[0]) &&
		!(rec1[1] >= rec2[3] || rec1[3] <= rec2[1])
}
