// 5179. 将二叉搜索树变平衡

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("res")
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先将原来的二叉搜索树转成数组，然后用二分法重建平衡树
func balanceBST(root *TreeNode) *TreeNode {
	var toSlice func(r *TreeNode) []int
	toSlice = func(r *TreeNode) []int {
		s := []int{}

		if r == nil {
			return s
		}
		s = append(s, r.Val)
		s = append(s, toSlice(r.Left)...)
		s = append(s, toSlice(r.Right)...)

		return s
	}

	s := toSlice(root)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	var toTree func(left, right int) *TreeNode

	toTree = func(left, right int) *TreeNode {
		if left >= right {
			return nil
		}
		mid := (left + right) / 2
		t := &TreeNode{
			Val: s[mid],
		}
		t.Left = toTree(left, mid)
		t.Right = toTree(mid+1, right)
		return t
	}

	ans := toTree(0, len(s))

	return ans
}
