package main

import (
	"fmt"
)

func main() {
	s := "1101"
	res := numSteps(s)
	fmt.Println("> ", res)
}

func numSteps(s string) int {
	step := 0

	for len(s) > 1 {
		n := len(s)
		last := s[n-1] - '0'
		step++

		if last == 0 {
			s = s[:n-1]
			// fmt.Println("0: ", s)
		} else {
			// fmt.Println("1: ", s)
			i := n - 2
			for i >= 0 {
				if s[i]-'0' == 0 {
					s = s[:i] + "1"
					step += n - len(s)
					// fmt.Println("1:::::", s, step, i, n, len(s))
					break
				}
				i--
			}

			if i < 0 {
				step += len(s)
				s = "1"
			}
		}
	}

	return step
}
