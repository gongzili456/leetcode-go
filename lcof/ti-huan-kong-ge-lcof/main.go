// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
// 面试题05. 替换空格
package main

import (
	"fmt"
)

func main() {
	s := "We are happy."
	res := replaceSpace(s)
	fmt.Println("res: ", res)
}

func replaceSpace(s string) string {
	ns := []byte{}

	for i := 0; i < len(s); i++ {
		u := s[i]
		if u == ' ' {
			ns = append(ns, '%', '2', '0')
		} else {
			ns = append(ns, u)
		}
	}
	return string(ns)
}
