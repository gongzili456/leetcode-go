// 5380. 数组中的字符串匹配
package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	res := stringMatching(words)
	fmt.Println("> ", res)
}

func stringMatching(words []string) []string {
	m := map[string]int{}
	ans := []string{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if strings.Contains(words[i], words[j]) {
				m[words[j]] = 1
			}
		}
	}

	for k := range m {
		ans = append(ans, k)
	}
	return ans
}
