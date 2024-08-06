package leetcode

import (
	"testing"
)

// 环形链表，leetcode 141
// https://leetcode.cn/problems/linked-list-cycle/description/
// 题目描述：
// 给你一个链表的头节点 head ，判断链表中是否有环。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。
//
// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
//
// 示例 2：
// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。
//
// 示例 3：
// 输入：head = [1], pos = -1
// 输出：false
// 解释：链表中没有环

func Test_HasCycle(t *testing.T) {

}

// hasCycle
// 思路：快慢指针p1、p2，只要有环，两个指针在某一时刻一定会相遇（类似于跑步中的套圈）
func hasCycle(head *ListNode) bool {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		// 两个指针相遇了，一定有环
		if p1 == p2 {
			return true
		}
	}
	return false
}
