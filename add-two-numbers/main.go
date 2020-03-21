// 两数相加 链表
// https://leetcode-cn.com/problems/add-two-numbers/
package main

import "fmt"

func main() {
	res := addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})

	for res != nil {
		fmt.Printf("%d", res.Val)
		res = res.Next
	}
}

// ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	ins := 0

	for l1 != nil || l2 != nil {
		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + ins
		ins = sum / 10

		p.Next = &ListNode{sum % 10, nil}
		p = p.Next
	}
	if ins == 1 {
		p.Next = &ListNode{1, nil}
		p = p.Next
	}
	return head.Next
}
