package 双指针解决链表问题

import "testing"

// 环形链表2 leetcode-142 mid
// 链接：https://leetcode.cn/problems/linked-list-cycle-ii/description/
// 题目介绍：
// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环
// 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）
// 如果 pos 是 -1，则在该链表中没有环
// 注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况
// 不允许修改 链表
//
// 示例1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点
//
// 示例2：
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
// 解释：链表中有一个环，其尾部连接到第一个节点
//
// 示例3
// 输入：head = [1], pos = -1
// 输出：返回 null
// 解释：链表中没有环

func Test_HasCycleV2(t *testing.T) {

}

// detectCycle hashmap去重，第一个重复的节点就是环的入口节点
func detectCycle(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{})
	// 直接遍历一遍链表，第一个重复的节点就是环的入口节点
	p := head
	for p != nil {
		if _, ok := nodeMap[p]; ok {
			return p
		} else {
			nodeMap[p] = struct{}{}
		}
		p = p.Next
	}
	return nil
}

// detectCycleV2 双指针解决链表问题
// 思路：
// p1和p2分别为慢指针和快指针，p1每次走一步，p2每次走2步
// 只要有环，p1和p2必定相遇
// 假设p1走了k步，那么p2一定走了p1的nk步（套圈），假设n=2，环的入口节点距相遇节点为m
// 那么头结点距离环的入口节点的距离就是：k-m（p1指针还没走完第一圈）
// 又因为p2指针从相遇节点开始，再走k步又可以会到相遇点，所以p1从相遇节点开始再走k-m步也可以回到环的入口节点
// 所以我们只需要再让一个指针p从头开始和p1一块走，当p和p1再次相遇时，p在的位置就是入口节点
func detectCycleV2(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			for p1 != p2 {
				p := head
				for p != p1 {
					p = p.Next
					p1 = p1.Next
				}
				return p
			}
		}
	}
	return nil
}
