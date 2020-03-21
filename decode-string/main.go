// https://leetcode-cn.com/problems/decode-string/
// 394. 字符串解码
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3[a2[c]]"

	res := decodeString(s)
	fmt.Println("res: ", res)
}

type element struct {
	multi int
	res   string
}

func decodeString(s string) string {
	stack := []element{}

	multi := ""
	res := ""
	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == '[' {
			n, _ := strconv.Atoi(multi)
			stack = append(stack, element{multi: n, res: res})
			multi = ""
			res = ""
			continue
		}

		if c == ']' {
			n := len(stack)
			pre := stack[n-1]
			stack = stack[:n-1]

			ts := ""
			for t := 0; t < pre.multi; t++ {
				ts += res
			}

			res = pre.res + ts
			multi = ""
			continue
		}

		if '0' <= c && c <= '9' {
			multi += string(c)
			continue
		}

		res += string(c)
	}

	return res
}
