package 双指针解决链表问题

import (
	"container/heap"
	"testing"
)

// 合并K个链表，leetcode-23 hard
// https://leetcode.cn/problems/merge-k-sorted-lists/description/
//
// 题目描述：
// 给你一个链表数组，每个链表都已经按升序排列
// 请你将所有链表合并到一个升序链表中，返回合并后的链表
//
// 示例1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//  1->4->5,
//  1->3->4,
//  2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
//
// 示例 2：
// 输入：lists = []
// 输出：[]
//
// 示例 3：
// 输入：lists = [[]]
// 输出：[]

func Test_MergeKLists(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Lists        []*ListNode
		ExpectResult *ListNode
	}{
		0: {
			Name: "case 1",
			Lists: []*ListNode{
				0: {
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
				1: {
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				2: {
					Val: 2,
					Next: &ListNode{
						Val:  6,
						Next: nil,
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
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		1: {
			Name:         "case 2",
			Lists:        make([]*ListNode, 0),
			ExpectResult: nil,
		},
		2: {
			Name:         "case 3",
			Lists:        make([]*ListNode, 0),
			ExpectResult: nil,
		},
	}

	for i, testCase := range testCaseList {
		lists := testCase.Lists
		expectResult := testCase.ExpectResult
		actualResult := mergeKListsV2(lists)
		for actualResult != nil && expectResult != nil {
			if actualResult.Val != expectResult.Val {
				t.Errorf("test case %v,not passed", i)
			}
			expectResult = expectResult.Next
			actualResult = actualResult.Next
		}
		t.Logf("test case %v,passed", i)
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists 合并K个有序链表，小顶堆思路，时间复杂度：O(n*Logk)，空间复杂度：O(k)
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	// 定义一个小顶堆
	minHeap := &hp{}
	// 将所有链表的头结点加入小顶堆
	for _, list := range lists {
		if list != nil {
			*minHeap = append(*minHeap, list)
		}
	}
	// 初始化小顶堆
	heap.Init(minHeap)

	// 开始构建最终的链表
	for minHeap.Len() > 0 {
		// 每次从小顶堆中取出一个最小值加入最终链表
		min := heap.Pop(minHeap).(*ListNode)
		if min.Next != nil {
			// 当前链表的下一个节点作为替补
			heap.Push(minHeap, min.Next)
		}
		cur.Next = min
		// 当前节点往后移动
		cur = cur.Next
	}

	return dummy.Next
}

// 实现heap接口方法，定义一个最小堆
type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(*ListNode))
}

func (h *hp) Pop() any {
	old := *h
	n := len(old)
	// 拿到堆顶的值，也就是此时堆的最小值
	min := old[n-1]
	// 删除堆顶
	*h = old[0 : n-1]
	// 返回最小堆
	return min
}

// mergeKListsV2 合并K个有序链表，分治思想，时间复杂度：O(n*Logk)，空间复杂度：O(logk)
func mergeKListsV2(lists []*ListNode) *ListNode {
	// base case
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	// 先合并左边的
	leftLists := mergeKListsV2(lists[0 : n/2])
	// 在合并右边的
	rightLists := mergeKListsV2(lists[n/2 : n])
	// 左右进行合并
	return merge2Lists(leftLists, rightLists)
}

// merge2Lists 合并两个有序链表
func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		// 找最小的最为cur的下一个节点
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		// cur指针往后走
		cur = cur.Next
	}
	// 剩下的链表直接插入到最后（因为有序）
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}
