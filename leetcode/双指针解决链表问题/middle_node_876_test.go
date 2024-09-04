package 双指针解决链表问题

import "testing"

// 链表的中间节点 leetcode 876
// https://leetcode.cn/problems/middle-of-the-linked-list/description/
//
// 题目描述：
// 给你单链表的头结点 head ，请你找出并返回链表的中间结点
// 如果有两个中间结点，则返回第二个中间结点
//
// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[3,4,5]
// 解释：链表只有一个中间结点，值为 3

// 示例 2：
// 输入：head = [1,2,3,4,5,6]
// 输出：[4,5,6]
// 解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点

func Test_MiddleNode(t *testing.T) {

}

// middleNode
// 思路：
// 快慢指针p1，p2，当快指针走完时，慢指针所在的位置就是中间的位置
func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1
}

// middleNodeV2 常规解法
// 把链表中的所有节点都塞到一个数组中，直接返回数组的中间元素即可
func middleNodeV2(head *ListNode) *ListNode {
	nodeList := make([]*ListNode, 0)
	tmp := head
	for tmp != nil {
		nodeList = append(nodeList, tmp)
		tmp = tmp.Next
	}

	return nodeList[len(nodeList)/2]
}
