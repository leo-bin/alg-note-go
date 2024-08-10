package leetcode

import (
	"testing"
)

// 相交链表，leetcode 160
// https://leetcode.cn/problems/intersection-of-two-linked-lists/
//
// 题目描述:
// 给你两个单链表的头节点 headA 和 headB
// 请你找出并返回两个单链表相交的起始节点,如两个链表不存在相交节点，返回null
// 题目数据保证整个链式结构中不存在环
//
// 示例1：
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
// 输出：Intersected at '8'
// 解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
// 从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
// 在A中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
// 请注意相交节点的值不为 1，因为在链表 A 和链表 B 之中值为 1 的节点 (A 中第二个节点和 B 中第三个节点) 是不同的节点。
// 换句话说，它们在内存中指向两个不同的位置
// 而链表 A 和链表 B 中值为 8 的节点 (A 中第三个节点，B 中第四个节点) 在内存中指向相同的位置。
// 注意，函数返回结果后，链表必须 保持其原始结构

func Test_GetIntersectionNode(t *testing.T) {

}

// getIntersectionNode hashmap,时间空间复杂度都是O(n)
func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{})
	// 1.遍历链表a，记录所有节点
	p1 := headA
	for p1 != nil {
		nodeMap[p1] = struct{}{}
		p1 = p1.Next
	}

	// 2.遍历链表b，判断是否有重复节点
	p2 := headB
	for p2 != nil {
		if _, ok := nodeMap[p2]; ok {
			return p2
		}
		p2 = p2.Next
	}

	return nil
}

// getIntersectionNodeV2 双指针技巧
// 思路：
// 先让p1和p2分别从两个链表的头结点开始遍历，遍历结束后，反过来从各自头结点开始在遍历一次
// 两个链表相交的话，两个指针走完的距离一定是相同的，并且一定会相遇（相遇后的节点都是重叠的）
func getIntersectionNodeV2(headA *ListNode, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		// p1走完了，从链表b从头开始再走一遍
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2走完了，从链表a从头开始再走一遍
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
