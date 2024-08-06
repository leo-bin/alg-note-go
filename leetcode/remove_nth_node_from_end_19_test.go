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

// removeNthFromEndV2 双指针思想（快慢指针）
// 指针p1和p2，p1先走n步（剩余路径就是k-n），然后p2和p1同时开始走，当p1走完时，p2就走了k-n步
// 此时p2所在的位置就是目标节点的位置，直接删除即可
// 注意，我们找到的是目标节点，但删除要知道目标节点的前置节点，所以要让p2少走一步
// 这里通过设置虚拟头结点来实现
func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	p1, p2 := head, dummy
	// 1.先让p1先走n步
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	// 2.p2开始跟着p1走，直到p1结束
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// 3.此时的p2就是目标节点的前置节点，直接删除即可
	p2.Next = p2.Next.Next

	return dummy.Next
}
