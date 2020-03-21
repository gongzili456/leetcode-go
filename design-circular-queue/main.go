// https://leetcode-cn.com/problems/design-circular-queue/
// 622. 设计循环队列
package main

import "fmt"

func main() {
	obj := Constructor(3)
	param1 := obj.EnQueue(1)
	param2 := obj.EnQueue(2)
	param3 := obj.EnQueue(3)
	fmt.Println(param1, param2, param3)

	param4 := obj.DeQueue()
	param5 := obj.Front()
	param6 := obj.Rear()
	fmt.Println(param4, param5, param6)

	param7 := obj.IsEmpty()
	param8 := obj.IsFull()
	fmt.Println(param7, param8)
}

// MyCircularQueue is
type MyCircularQueue struct {
	Cap  int
	Q    []int
	Tail int
	Head int
}

// Constructor Initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Cap:  k,
		Q:    make([]int, k),
		Tail: -1,
		Head: -1,
	}
}

// EnQueue Insert an element into the circular queue. Return true if the operation is successful.
func (c *MyCircularQueue) EnQueue(value int) bool {
	fmt.Println("Q: ", c.Q, c.Head, c.Tail)
	if c.IsFull() {
		return false
	}
	if c.IsEmpty() {
		c.Head = 0
	}

	// if c.Tail == c.Cap-1 {
	// 	c.Tail = 0
	// } else {
	// 	c.Tail++
	// }
	c.Tail = (c.Tail + 1) % c.Cap
	c.Q[c.Tail] = value
	return true
}

// DeQueue Delete an element from the circular queue. Return true if the operation is successful.
func (c *MyCircularQueue) DeQueue() bool {
	if c.IsEmpty() {
		return false
	}

	if c.Head == c.Tail {
		c.Head = -1
		c.Tail = -1
		return true
	}

	// if c.Head == c.Cap-1 {
	// 	c.Head = 0
	// } else {
	// 	c.Head++
	// }

	c.Head = (c.Head + 1) % c.Cap

	return true
}

// Front Get the front item from the queue.
func (c *MyCircularQueue) Front() int {
	if c.IsEmpty() {
		return -1
	}

	return c.Q[c.Head]
}

// Rear Get the last item from the queue.
func (c *MyCircularQueue) Rear() int {
	if c.IsEmpty() {
		return -1
	}

	return c.Q[c.Tail]
}

// IsEmpty Checks whether the circular queue is empty or not.
func (c *MyCircularQueue) IsEmpty() bool {
	if c.Head == -1 {
		return true
	}
	return false
}

// IsFull Checks whether the circular queue is full or not.
func (c *MyCircularQueue) IsFull() bool {
	if c.IsEmpty() {
		return false
	}

	// if c.Head > c.Tail && c.Head-c.Tail == 1 {
	// 	return true
	// }

	// if c.Head < c.Tail && c.Tail-c.Head == c.Cap-1 {
	// 	return true
	// }

	return (c.Tail+1)%c.Cap == c.Head
}
