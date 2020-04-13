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

// 深度搜索，如果一个节点的所有子节点（含节点本身）包含所有目标节点，那么该节点就是最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 给每个节点分配权重，搜索目标节点是1，其余为0，
	// 深度搜索，将权重的和返回上层，如果权重和 >=2 则发现公共祖先节点
	var ans *TreeNode
	var search func(r *TreeNode) bool
	search = func(node *TreeNode) bool {
		if ans != nil {
			return false
		}

		if node == nil {
			return false
		}

		var l, r, c int

		if left := search(node.Left); left {
			l = 1
		}
		if right := search(node.Right); right {
			r = 1
		}
		if node == p || node == q {
			c = 1
		}

		if l+r+c >= 2 {
			ans = node
		}

		return (l + r + c) > 0
	}

	search(root)
	return ans
}
