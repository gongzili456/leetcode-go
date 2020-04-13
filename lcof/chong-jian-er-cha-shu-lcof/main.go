// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
// 面试题07. 重建二叉树
package main

import "fmt"

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	res := buildTree(preorder, inorder)
	fmt.Println("res: ", res)
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

	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		root := &TreeNode{
			Val: preorder[rootIndex],
		}
		rootIndex++

		mi := m[root.Val]

		root.Left = build(left, mi-1)
		root.Right = build(mi+1, right)

		return root
	}

	return build(rootIndex, len(preorder)-1)
}
