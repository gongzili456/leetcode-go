// https://leetcode-cn.com/problems/min-stack/
// 最小栈
package main

import (
	"fmt"
	"math"
)

func main() {
	obj := Constructor()

	obj.Push(5)
	obj.Push(3)
	obj.Push(2)
	obj.Push(4)

	obj.Pop()

	param3 := obj.Top()
	param4 := obj.GetMin()

	fmt.Println(param3, param4)
}

// StackNode is
type StackNode struct {
	Val  int
	Min  int
	Next *StackNode
}

// MinStack is
type MinStack struct {
	head *StackNode
}

// Constructor initialize your data structure here
func Constructor() MinStack {
	return MinStack{&StackNode{-1, math.MinInt32, nil}}
}

// Push is
func (m *MinStack) Push(x int) {
	curr := m.head.Next

	// link is nil
	if curr == nil {
		m.head.Next = &StackNode{Val: x, Min: x, Next: nil}
		return
	}

	var newbie *StackNode
	if curr.Min < x {
		newbie = &StackNode{Val: x, Min: curr.Min, Next: curr}
	} else {
		newbie = &StackNode{Val: x, Min: x, Next: curr}
	}

	m.head.Next = newbie
}

// Pop is
func (m *MinStack) Pop() {
	m.head.Next = m.head.Next.Next
}

// Top is
func (m *MinStack) Top() int {
	return m.head.Next.Val
}

// GetMin is
func (m *MinStack) GetMin() int {
	return m.head.Next.Min
}
