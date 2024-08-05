package leetcode

import "testing"

// 分隔链表 leetcode 86
// https://leetcode.cn/problems/partition-list/
// 题目介绍：
// 给你一个链表的头节点head和一个特定值x，请你对链表进行分隔，使得所有小于x的节点都出现在大于或等于x的节点之前
// 你应当保留两个分区中每个节点的初始相对位置
//
// 示例1：
// 输入：head = [1,4,3,2,5,2], x = 3
// 输出：[1,2,2,4,3,5]
//
// 示例 2：
// 输入：head = [2,1], x = 2
// 输出：[1,2]
func Test_ListPartition(t *testing.T) {
	testCaseList := []struct {
		Head         *ListNode
		X            int
		ExpectResult *ListNode
	}{
		0: {
			Head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  2,
									Next: nil,
								},
							},
						},
					},
				},
			},
			X: 3,
			ExpectResult: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for i, testCase := range testCaseList {
		head := testCase.Head
		x := testCase.X
		expectResult := testCase.ExpectResult
		actualResult := listPartition(head, x)
		if actualResult == nil {
			t.Errorf("excute err")
		}
		e := expectResult
		a := actualResult
		for e != nil && a != nil {
			if e.Val != a.Val {
				t.Errorf("test case %v,not passed", i)
			}
			e = e.Next
			a = a.Next
		}
		t.Logf("test case %v,passed", i)
	}
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func listPartition(head *ListNode, x int) *ListNode {
	leftDummy := &ListNode{}
	rightDummy := &ListNode{}
	left := leftDummy
	right := rightDummy

	for head != nil {
		// 找比x要小的
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			// 找大于或等于x的
			right.Next = head
			right = right.Next
		}
		// head指针往后移
		head = head.Next
	}
	// 清除右链表的最后一个节点指向，防止成环
	right.Next = nil
	// 连接左右链表
	left.Next = rightDummy.Next

	return leftDummy.Next
}
