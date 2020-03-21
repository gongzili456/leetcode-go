// https://leetcode-cn.com/problems/implement-stack-using-queues/
// 225. 用队列实现栈
package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)

	param2 := obj.Pop()
	param3 := obj.Top()
	param4 := obj.Empty()

	fmt.Println("res: ", param2, param3, param4)
	fmt.Println("obj: ", obj)
}

// MyStack is
type MyStack struct {
	Q []int
}

// Constructor Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		Q: make([]int, 0),
	}
}

// Push element x onto stack. */
func (s *MyStack) Push(x int) {
	s.Q = append(s.Q, x)
}

// Pop Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	if s.Empty() {
		return -1
	}
	n := len(s.Q)
	x := s.Q[n-1]
	s.Q = s.Q[:n-1]
	return x
}

// Top Get the top element. */
func (s *MyStack) Top() int {
	if s.Empty() {
		return -1
	}

	return s.Q[len(s.Q)-1]
}

// Empty Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.Q) <= 0
}
