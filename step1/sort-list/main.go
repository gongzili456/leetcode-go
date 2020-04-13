// https://leetcode-cn.com/problems/sort-list/
// 链表排序
package main

import "fmt"

func main() {
	res := sortList(&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}})

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

// 迭代 失败
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := length(head)

	P := &ListNode{0, nil}
	P.Next = head

	for i := 1; i < length; i <<= 1 {
		prev := P
		curr := P.Next

		for curr != nil {
			left := curr
			right := split(left, i)
			curr = split(right, i)
			prev.Next = merge(left, right)

			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}

	return P.Next
}

func split(head *ListNode, step int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for i := 0; i < step && head.Next != nil; i++ {
		head = head.Next
	}

	right := head.Next
	head.Next = nil

	return right
}

func merge(l1, l2 *ListNode) *ListNode {
	P := &ListNode{0, nil}
	H := P

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			H.Next = l1
			l1 = l1.Next
		} else {
			H.Next = l2
			l2 = l2.Next
		}

		H = H.Next
	}

	if l1.Next != nil {
		H.Next = l1
	} else {
		H.Next = l2
	}

	return P.Next
}

func length(head *ListNode) int {
	p := head

	length := 0
	for p != nil {
		length++
		p = p.Next
	}
	return length
}

// 递归
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 寻找中间点
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	left := sortList(slow.Next)
	slow.Next = nil
	right := sortList(head)

	p := &ListNode{0, nil}
	H := p

	// sort
	for left != nil && right != nil {
		lv := left.Val
		rv := right.Val

		if lv < rv {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}

	if left != nil {
		p.Next = left
	} else {
		p.Next = right
	}

	return H.Next
}
