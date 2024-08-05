package leetcode

import "testing"

// 删除链表倒数第N个节点，leetcode 19
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
//
// 题目描述：
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
//
// 示例1：
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]
//
// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]

func Test_RemoveNthFromEnd(t *testing.T) {

}

// removeNthFromEnd 遍历写法，删除链表的倒数第N个节点，思路：倒数第n个节点，就是正数第k-n+1个节点，k是节点总长度
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1.遍历1遍链表，得到链表总长度k
	k := 0
	p1 := head
	for p1 != nil {
		k++
		// 移动p1指针
		p1 = p1.Next
	}

	// 2.开始找倒数第n+1个节点
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < k-n; i++ {
		cur = cur.Next
	}
	// 删除倒数第n个节点
	cur.Next = cur.Next.Next

	return dummy.Next
}
