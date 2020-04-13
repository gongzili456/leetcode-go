// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// 150. 逆波兰表达式求值
package main

import (
	"fmt"
	"strconv"
)

func main() {
	polish := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	res := evalRPN(polish)
	fmt.Println("res: ", res)
}

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	op := map[string]string{"+": "+", "-": "-", "*": "*", "/": "/"}

	var pop func() int
	pop = func() int {
		n := len(stack)
		x := stack[n-1]
		stack = stack[:n-1]
		i, _ := strconv.Atoi(x)
		return i
	}

	for i := 0; i < len(tokens); i++ {
		x := tokens[i]
		if o, ok := op[x]; ok {
			one, tow := pop(), pop()
			res := 0
			switch o {
			case "+":
				res = tow + one
			case "-":
				res = tow - one
			case "*":
				res = tow * one
			case "/":
				res = tow / one
			}
			stack = append(stack, strconv.Itoa(res))
		} else {
			stack = append(stack, x)
		}
	}

	return pop()
}
