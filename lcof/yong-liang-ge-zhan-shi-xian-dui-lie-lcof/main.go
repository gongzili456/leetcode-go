// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
// 面试题09. 用两个栈实现队列
package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println("> ", obj.DeleteHead())

	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Println("> ", obj)

	fmt.Println("> ", obj.DeleteHead())
	fmt.Println("> ", obj)

	fmt.Println("> ", obj.DeleteHead())
}

// CQueue is
type CQueue struct {
	s1 []int
	s2 []int
}

// Constructor is
func Constructor() CQueue {
	return CQueue{s1: []int{}, s2: []int{}}
}

// AppendTail is
func (c *CQueue) AppendTail(value int) {
	for len(c.s2) > 0 {
		n := len(c.s2)
		s := c.s2[n-1]
		c.s2 = c.s2[0 : n-1]
		c.s1 = append(c.s1, s)
	}
	c.s1 = append(c.s1, value)
}

// DeleteHead is
func (c *CQueue) DeleteHead() int {
	if len(c.s2) == 0 && len(c.s1) == 0 {
		return -1
	}

	for len(c.s1) > 0 {
		n := len(c.s1)
		s := c.s1[n-1]
		c.s1 = c.s1[0 : n-1]
		c.s2 = append(c.s2, s)
	}

	n := len(c.s2)
	s := c.s2[n-1]
	c.s2 = c.s2[0 : n-1]
	return s
}
