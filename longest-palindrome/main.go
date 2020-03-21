// https://leetcode-cn.com/problems/longest-palindrome/
// 409. 最长回文串
package main

import "fmt"

func main() {
	s := "abccccdd"
	res := longestPalindrome(s)
	fmt.Println("res: ", res)
}

// 只能有一个奇数字符，把所有奇数个数的取余，即把奇数变成偶数剩余的奇数长度
func longestPalindrome(s string) int {
	m := map[rune]int{}

	for _, c := range s {
		if v, ok := m[c]; ok {
			m[c] = v + 1
		} else {
			m[c] = 1
		}
	}

	count := 0
	for _, v := range m {
		count += v % 2
	}

	if count > 0 {
		// 去掉所有剩余奇数长度，然后 +1 个奇数长度
		return len(s) - count + 1
	}
	return len(s)
}
