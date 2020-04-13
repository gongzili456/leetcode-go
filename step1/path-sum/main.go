// https://leetcode-cn.com/problems/path-sum/
// 112. 路径总和
package main

func main() {}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自顶至下
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	has := false

	var pathSum func(r *TreeNode, s int)
	pathSum = func(r *TreeNode, s int) {
		s += r.Val
		if r.Left == nil && r.Right == nil && sum == s {
			has = true
		}

		if r.Left != nil {
			pathSum(r.Left, s)
		}

		if r.Right != nil {
			pathSum(r.Right, s)
		}
	}

	pathSum(root, 0)

	return has
}
