// https://leetcode-cn.com/problems/multiply-strings/solution/you-hua-ban-shu-shi-da-bai-994-by-breezean/
// multiply-strings
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num1 := "9"
	num2 := "9"
	res := multiply(num1, num2)
	fmt.Println("res: ", res)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := "0"

	var multipl func(x string, y int, m int) string
	multipl = func(x string, y, m int) string {
		fmt.Println("> ", x, y, m)
		s := make([]string, len(x))
		a := 0
		for i := len(x) - 1; i >= 0; i-- {
			r := int(x[i]-'0')*y + a
			a = r / 10
			s[i] = strconv.Itoa(r % 10)
		}

		for m > 0 {
			s = append(s, "0")
			m--
		}

		if a > 0 {
			return strconv.Itoa(a) + strings.Join(s, "")
		}

		return strings.Join(s, "")
	}

	var add func(x, y string) string
	add = func(x, y string) string {
		fmt.Println("add: ", x, y)
		if len(x) < len(y) {
			x, y = y, x
		}

		s := make([]string, len(x))
		a := 0
		for i, j := len(x)-1, len(y)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
			var xv, yv int
			if i >= 0 {
				xv = int(x[i] - '0')
			}
			if j >= 0 {
				yv = int(y[j] - '0')
			}
			sum := xv + yv + a
			a = sum / 10
			s[i] = strconv.Itoa(sum % 10)
		}

		if a > 0 {
			return strconv.Itoa(a) + strings.Join(s, "")
		}

		return strings.Join(s, "")
	}

	for i := len(num2) - 1; i >= 0; i-- {
		m := multipl(num1, int(num2[i]-'0'), len(num2)-i-1)
		fmt.Println("m: ", m)
		result = add(result, m)
	}

	return result
}
