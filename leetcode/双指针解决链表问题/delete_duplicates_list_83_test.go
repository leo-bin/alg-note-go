package 双指针解决链表问题

import "testing"

// 删除排序链表中的重复元素 leetcode-83
// 链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// 题目描述：
// 给定一个已排序的链表的头 head
// 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表
//
// 示例1：
// 输入：head = [1,1,2]
// 输出：[1,2]
//
// 示例2：
// 输入：head = [1,1,2,3,3]
// 输出：[1,2,3]

func Test_DeleteDuplicatesList(t *testing.T) {

}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// deleteDuplicates hashmap去重 时间O（n），空间O（n）
func deleteDuplicates(head *ListNode) *ListNode {
	// base case
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	nodeSet := make(map[int]struct{})
	p := head
	newHead := &ListNode{}
	tmp := newHead
	for p != nil {
		if _, exits := nodeSet[p.Val]; !exits {
			nodeSet[p.Val] = struct{}{}
			tmp.Next = &ListNode{Val: p.Val, Next: nil}
			tmp = tmp.Next
		}
		p = p.Next
	}

	return newHead.Next
}

// deleteDuplicates 快慢指针原地修改链表解法，时间O（n），空间O（1）
// 思想：
// 1.慢指针修改链表的关系，快指针在前面探路，查找重复的元素
// 2.slow和fast同时出发，当遇到连续重复的元素时，fast继续往前走，slow不动
// 3.当出现和slow不同的元素时（此时找到了下一个不重复的元素），需要修改链表的关系，将slow指向fast，并断开slow原来的关系
func deleteDuplicatesV2(head *ListNode) *ListNode {
	// base case
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = fast
		}
		fast = fast.Next
	}
	// 最后一个节点如果也是连续的，需要去重
	slow.Next = nil

	return head
}
