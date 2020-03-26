// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 105. 从前序与中序遍历序列构造二叉树
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	rootIndex := 0
	m := map[int]int{}
	for i, v := range inorder {
		m[v] = i
	}

	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		root := preorder[rootIndex]
		rootIndex++

		mi := m[root]

		leftTree := helper(left, mi-1)
		rightTree := helper(mi+1, right)

		return &TreeNode{Val: root, Left: leftTree, Right: rightTree}
	}
	return helper(rootIndex, len(preorder)-1)
}
