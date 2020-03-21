// https://leetcode-cn.com/problems/implement-queue-using-stacks
// 用栈实现队列
package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)

	param3 := obj.Peek()  // 1
	param2 := obj.Pop()   // 1
	param4 := obj.Empty() // false

	fmt.Println("> ", param2, param3, param4)
}

// MyQueue is
type MyQueue struct {
	Stack []int
}

// Constructor Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Stack: make([]int, 0),
	}
}

// Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.Stack = append(q.Stack, x)
}

// Pop Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.Empty() {
		return -1
	}
	x := q.Stack[0]
	q.Stack = q.Stack[1:]
	return x
}

// Peek Get the front element. */
func (q *MyQueue) Peek() int {
	if q.Empty() {
		return -1
	}
	return q.Stack[0]
}

// Empty Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.Stack) <= 0
}
