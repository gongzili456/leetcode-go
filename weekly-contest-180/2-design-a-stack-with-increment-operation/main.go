package main

import "fmt"

func main() {
	obj := Constructor(10)
	obj.Push(1)
	param2 := obj.Pop()

	obj.Push(1)
	obj.Push(12)
	obj.Push(4)
	obj.Push(8)
	obj.Push(4)

	obj.Increment(10, 10)

	fmt.Println("res: ", param2)
	fmt.Print("SS: ", obj)
}

// CustomStack is
type CustomStack struct {
	maxSize int
	stack   []int
}

// Constructor is
func Constructor(maxSize int) CustomStack {
	return CustomStack{maxSize, []int{}}
}

// Push is
func (c *CustomStack) Push(x int) {
	if len(c.stack) >= c.maxSize {
		return
	}
	c.stack = append(c.stack, x)
}

// Pop is
func (c *CustomStack) Pop() int {
	if len(c.stack) <= 0 {
		return -1
	}
	x := c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
	return x
}

// Increment is
func (c *CustomStack) Increment(k int, val int) {
	for i := 0; i < len(c.stack) && i < c.maxSize && i < k; i++ {
		c.stack[i] += val
	}
}
