// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 面试题06. 从尾到头打印链表
package main

func main() {

}

// ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	ans := []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
	}

	return ans
}
