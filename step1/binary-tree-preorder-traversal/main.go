// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
// 144. 二叉树的前序遍历
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
	res := preorderTraversal(t)
	fmt.Println("res: ", res)
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	stack = append(stack, root)

	ans := []int{}

	for len(stack) > 0 {
		n := len(stack)
		r := stack[n-1]
		stack = stack[:n-1]

		if r == nil {
			continue
		}

		ans = append(ans, r.Val)

		stack = append(stack, r.Right, r.Left)
	}

	return ans
}
