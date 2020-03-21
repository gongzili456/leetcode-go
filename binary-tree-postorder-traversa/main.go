// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
// 145. 二叉树的后序遍历
package main

import "fmt"

func main() {
	t := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}

	res := postorderTraversal(t)
	fmt.Println("res: ", res)
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序：   根，左，右
// 反向前序：跟，右，左
// 反转结果：左，右，跟
// left right center
func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	stack = append(stack, root)

	ans := []int{}

	type node struct {
		Val  int
		Next *node
	}

	p := &node{}

	for len(stack) > 0 {
		n := len(stack)
		curr := stack[n-1]
		stack = stack[:n-1]

		if curr == nil {
			continue
		}

		// 插入链表的 头部
		c := &node{Val: curr.Val, Next: p.Next}
		p.Next = c

		// 反向前序遍历
		stack = append(stack, curr.Left, curr.Right)
	}

	for p.Next != nil {
		p = p.Next
		ans = append(ans, p.Val)
	}

	return ans
}

// 写法二
// func postorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	stack := []*TreeNode{}

// 	type node struct {
// 		Val  int
// 		Next *node
// 	}

// 	p := &node{}

// 	curr := root
// 	for curr != nil || len(stack) > 0 {
// 		for curr != nil {
// 			c := &node{Val: curr.Val, Next: p.Next}
// 			p.Next = c

// 			stack = append(stack, curr)
// 			curr = curr.Right
// 		}

// 		n := len(stack)
// 		curr = stack[n-1]
// 		stack = stack[:n-1]

// 		curr = curr.Left
// 	}

// 	for p.Next != nil {
// 		p = p.Next
// 		ans = append(ans, p.Val)
// 	}
// 	return ans
// }
