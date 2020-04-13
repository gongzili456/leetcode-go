// https://leetcode-cn.com/problems/valid-parentheses/
// 20. 有效的括号
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "(()[]{}"

	res := isValid(s)
	fmt.Println("res: ", res)
}

func isValid(s string) bool {
	stack := []byte{}
	char := []byte{'(', ')', '{', '}', '[', ']'}

	for i := 0; i < len(s); i++ {
		c := s[i]

		index := bytes.IndexByte(char, c)

		if index%2 == 0 {
			stack = append(stack, c)
			continue
		}

		n := len(stack)
		if n == 0 {
			return false
		}

		left := stack[n-1]
		stack = stack[:n-1]

		if bytes.IndexByte(char, left)+1 != index {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
