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

// 反转链表，206，easy
// https://leetcode.cn/problems/reverse-linked-list/
// 思路：
// 1.迭代：用三个指针记录下pre、cur、next的位置，原地进行反转
// 2.递归：head.next.next = head

// 迭代
func reverseListIter(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// 定义三个指针
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head.Next
	for cur != nil {
		// 反转
		cur.Next = pre
		// 移动pre,cur和next
		pre = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}
	return pre
}

// 递归: head.Next.Next = head
func reverseListDfs(head *ListNode) *ListNode {
	// 递归结束条件
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListDfs(head.Next)
	// 在这里原地反转
	head.Next.Next = head
	head.Next = nil
	return newHead
}
