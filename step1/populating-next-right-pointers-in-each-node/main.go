// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
// 116. 填充每个节点的下一个右侧节点指针
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
// 117. 填充每个节点的下一个右侧节点指针 II
package main

func main() {}

// Node is
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Root.Left.Next = Root.Right
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// 逐层遍历
	q := []*Node{}
	q = append(q, root)

	for len(q) > 0 {
		w := len(q)
		for i := 0; i < w; i++ {
			n := q[i]

			if i < w-1 {
				n.Next = q[i+1]
			}
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		q = q[w:]
	}

	return root
}
