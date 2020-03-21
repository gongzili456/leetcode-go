// 5352. 生成每种字符都是奇数个的字符串
package main

import "fmt"

func main() {
	res := generateTheString(4)

	fmt.Println("res: ", res)
}

func generateTheString(n int) string {

	// 偶数
	if n%2 == 0 {
		return getStr(0, n-1) + getStr(1, 1)
	} else {
		return getStr(0, n)
	}
}

func getStr(offset, n int) string {
	chars := []byte{}

	for i := 0; i < 26; i++ {
		chars = append(chars, byte('a'+i))
	}

	str := ""
	for i := 0; i < n; i++ {
		str += string(chars[offset])
	}
	return string(str)
}
