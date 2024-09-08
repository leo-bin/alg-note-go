package 双指针解决链表问题

import "testing"

// 合并两个有序链表，leetcode-21 easy
// https://leetcode.cn/problems/merge-two-sorted-lists/
// 题目介绍：
// 将两个升序链表合并为一个新的 升序 链表并返回
// 新链表是通过拼接给定的两个链表的所有节点组成的
//
// 示例 1：
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
// 示例 2：
// 输入：l1 = [], l2 = []
// 输出：[]
//
// 示例 3：
// 输入：l1 = [], l2 = [0]
// 输出：[0]

func Test_MergeTwoLists(t *testing.T) {
	testCaseList := []struct {
		List1        *ListNode
		List2        *ListNode
		ExpectResult *ListNode
	}{
		0: {
			List1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			List2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			ExpectResult: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
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
		list1 := testCase.List1
		list2 := testCase.List2
		expectResult := testCase.ExpectResult
		actualResult := mergeTwoLists(list1, list2)
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		// 移动cur指针
		cur = cur.Next
	}
	// 最后剩下的节点直接放到最后拼接起来
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}
