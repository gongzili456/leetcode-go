// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 二叉树的最近公共祖先
package main

import "fmt"

func main() {

	root := newTreeNode(3)
	q := newTreeNode(4)

	p := &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, q}}

	root.Left = p
	root.Right = &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}}

	res := lowestCommonAncestor(root, p, q)
	fmt.Println("res: ", res)

}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var ans *TreeNode
	var search func(r *TreeNode) bool

	search = func(r *TreeNode) bool {
		if r == nil {
			return false
		}

		var self, left, right int

		if l := search(r.Left); l {
			left = 1
		}

		if r := search(r.Right); r {
			right = 1
		}

		if r == p || r == q {
			self = 1
		}
		fmt.Println("::: ", left+self+right)
		if (left + self + right) >= 2 {
			ans = r
		}

		return (left + self + right) > 0
	}

	search(root)
	return ans
}
