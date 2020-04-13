// 5353. 灯泡开关 III
package main

import (
	"fmt"
)

func main() {
	res := numTimesAllBlue([]int{2, 1, 3, 5, 4})
	fmt.Println("res: ", res)
}

func numTimesAllBlue(light []int) int {
	count := 0
	max := 0
	for i, v := range light {
		if max < v {
			max = v
		}

		if max == i+1 {
			count++
		}
	}
	return count
}
