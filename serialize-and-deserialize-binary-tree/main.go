// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
// 297. 二叉树的序列化与反序列化
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	obj := Constructor()
	data := obj.serialize(root)
	ans := obj.deserialize(data)

	fmt.Println("serialize: ", data)
	fmt.Println()

	q := []*TreeNode{ans}

	for len(q) > 0 {
		r := q[0]
		q = q[1:]
		if r == nil {
			fmt.Printf("%s,", "null")
			continue
		}
		fmt.Printf("%d,", r.Val)

		q = append(q, r.Left, r.Right)
	}
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec is
type Codec struct {
}

// Constructor is
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	// 前序深度遍历：跟左右
	stack := []*TreeNode{root}
	ans := []string{}

	for len(stack) > 0 {
		n := len(stack)
		r := stack[n-1]
		stack = stack[0 : n-1]

		if r == nil {
			ans = append(ans, "null")
			continue
		}

		ans = append(ans, fmt.Sprint(r.Val))
		stack = append(stack, r.Right, r.Left)
	}
	return "[" + strings.Join(ans, ",") + "]"
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	// [1,2,null,null,3,4,null,null,5,null,null]
	// 前序深度遍历：跟左右
	data = data[1 : len(data)-1]
	fmt.Println("d: ", data)

	ds := strings.Split(data, ",")
	fmt.Println("ds: ", ds)

	var des func() *TreeNode
	des = func() *TreeNode {
		d := ds[0]
		ds = ds[1:]
		if d == "null" {
			return nil
		}

		n, _ := strconv.Atoi(d)
		r := &TreeNode{Val: n}
		r.Left = des()
		r.Right = des()
		return r
	}

	return des()
}
