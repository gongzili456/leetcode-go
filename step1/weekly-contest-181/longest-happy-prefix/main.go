// https://leetcode-cn.com/contest/weekly-contest-181/problems/longest-happy-prefix/
// 5367. 最长快乐前缀
package main

import "fmt"

func main() {
	// s := "level"
	// s := "leetcodeleet"
	s := "ababab"

	res := longestPrefix(s)
	fmt.Println("res: ", res)
}

func longestPrefix(s string) string {
	ans := ""
	n := len(s) - 1

	l, r := 0, n

	for l < n && r > 0 {
		ls := s[0 : l+1]
		rs := s[r:]

		if ls == rs && len(ls) > len(ans) {
			ans = ls
		}
		l++
		r--
	}

	return ans
}
