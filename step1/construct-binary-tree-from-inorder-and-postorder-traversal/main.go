// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 106. 从中序与后序遍历序列构造二叉树
package main

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	buildTree(inorder, postorder)
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序：左右中
// 中序：左中右
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 首先，找到根节点，后序序列中的最后一个，永远是树的根节点
	// 然后处理右子树
	// 处理左子树
	// 递归重复
	rootIndex := len(postorder) - 1
	m := map[int]int{}
	for i, v := range inorder {
		m[v] = i
	}

	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// 确定子树的跟节点
		root := postorder[rootIndex]

		// 确定下一轮的根节点位置
		rootIndex--

		// 找到跟节点在中序序列的位置，其左侧为左子树，右侧为右子树
		mi := m[root]

		// 注意要先右后左
		rightNode := helper(mi+1, right)
		leftNode := helper(left, mi-1)

		return &TreeNode{Val: root, Left: leftNode, Right: rightNode}
	}

	return helper(0, rootIndex)
}
