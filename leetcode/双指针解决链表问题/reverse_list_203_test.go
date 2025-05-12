package 双指针解决链表问题

import "testing"

// 反转链表，easy，203
// https://leetcode.cn/problems/reverse-linked-list/
// 给你单链表的头节点 head
// 请你反转链表，并返回反转后的链表
//
// 示例1:
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
//
// 示例2:
// 输入：head = [1,2]
// 输出：[2,1]

func TestReverseList(t *testing.T) {

}

// 递归法
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	return newHead
}
