// https://leetcode-cn.com/problems/same-tree/
// 100. 相同的树
package main

import "fmt"

func main() {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{Val: 3},
	}

	res := isSameTree(t1, t1)
	fmt.Println("res: ", res)
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	q1, q2 := []*TreeNode{}, []*TreeNode{}

	q1 = append(q1, p)
	q2 = append(q2, q)

	for len(q1) > 0 || len(q2) > 0 {
		var pp, qq *TreeNode
		if len(q1) > 0 {
			pp = q1[0]
			q1 = q1[1:]
		}
		if len(q2) > 0 {
			qq = q2[0]
			q2 = q2[1:]
		}

		if pp == nil && qq == nil {
			continue
		}

		if pp == nil {
			return false
		}

		if qq == nil {
			return false
		}

		if pp.Val != qq.Val {
			return false
		}

		if pp != nil {
			q1 = append(q1, pp.Left, pp.Right)
		}

		if qq != nil {
			q2 = append(q2, qq.Left, qq.Right)
		}
	}

	return true
}
