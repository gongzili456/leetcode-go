// https://leetcode-cn.com/problems/compress-string-lcci/
// 面试题 01.06. 字符串压缩
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "aabcccccaaa"

	res := compressString(s)
	fmt.Println("res: ", res)
}

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}

	s := []string{}
	c := 1

	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			c++
			continue
		}
		s = append(s, string(S[i-1]), strconv.Itoa(c))
		c = 1
	}

	s = append(s, string(S[len(S)-1]), strconv.Itoa(c))

	if len(S) <= len(s) {
		return S
	}
	return strings.Join(s, "")
}
