package leetcode

import "testing"

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

// detectCycleV2 双指针
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
