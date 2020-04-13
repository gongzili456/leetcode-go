// https://leetcode-cn.com/problems/open-the-lock/solution/da-kai-zhuan-pan-suo-by-leetcode/
// 打开转盘锁
package main

import (
	"fmt"
)

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"

	res := openLock(deadends, target)
	fmt.Println("res: ", res)
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return -1
	}

	dead := map[string]int{}

	for _, v := range deadends {
		dead[v] = 0
	}

	if _, ok := dead["0000"]; ok {
		return -1
	}

	q := []string{"0000"}
	dead["0000"] = 0

	for len(q) != 0 {
		s := q[0]
		q = q[1:]
		depth := dead[s]

		// 对4位置的数字分别进行 +1， -1 的操作
		for i := 0; i < 4; i++ {
			for d := -1; d <= 1; d += 2 {
				x := (int(s[i]-'0')+d+10)%10 + '0' // 注意 byte 和 int 之间的转换
				n := s[0:i] + string(x) + s[i+1:]

				if _, ok := dead[n]; ok {
					continue
				}

				dead[n] = depth + 1
				q = append(q, n)

				if n == target {
					return dead[n]
				}
			}
		}

	}

	return -1
}
