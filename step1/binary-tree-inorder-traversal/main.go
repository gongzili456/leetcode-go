// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-by-leetcode/
// 94. 二叉树的中序遍历
package main

import "fmt"

func main() {
	t := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res := inorderTraversal(t)
	fmt.Println("res: ", res)
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序：跟左右
// 后序：左右跟
// 中序：左跟右
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	ans := []int{}

	var pop func() *TreeNode
	pop = func() *TreeNode {
		n := len(stack)
		x := stack[n-1]
		stack = stack[:n-1]
		return x
	}

	curr := root
	for curr != nil || len(stack) > 0 {
		// 将所有 left 加入栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Left or center Node
		curr = pop()
		ans = append(ans, curr.Val)

		// Push right nodes
		curr = curr.Right
	}

	return ans
}
