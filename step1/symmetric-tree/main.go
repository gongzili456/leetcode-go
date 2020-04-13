// https://leetcode-cn.com/problems/symmetric-tree/solution/dui-cheng-er-cha-shu-by-leetcode/
// 101. 对称二叉树
package main

// 如果同时满足下面的条件，两个树互为镜像：
//
// 它们的两个根结点具有相同的值。
// 每个树的右子树都与另一个树的左子树镜像对称。
func main() {}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法
func isSymmetric2(root *TreeNode) bool {
	Q := []*TreeNode{}
	Q = append(Q, root, root)

	for len(Q) > 0 {
		t1, t2 := Q[0], Q[1]
		Q = Q[2:]

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}

		// 注意顺序是对称的
		Q = append(Q, t1.Left, t2.Right, t1.Right, t2.Left)
	}

	return true
}

// 递归法
func isSymmetric(root *TreeNode) bool {
	var isMirror func(t1, t2 *TreeNode) bool
	isMirror = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}

		if t1 == nil || t2 == nil {
			return false
		}

		return t1.Val == t2.Val &&
			isMirror(t1.Left, t2.Right) &&
			isMirror(t1.Right, t2.Left)
	}
	return isMirror(root, root)
}
