package 双指针解决链表问题

import "testing"

// 回文链表，234,easy
// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表
// 如果是，返回 true ；否则，返回 false
// https://leetcode.cn/problems/palindrome-linked-list/description/
//
// 示例1:
// 输入：head = [1,2,2,1]
// 输出：true
func TestIsPalindrome(t *testing.T) {

}

// isPalindrome
// 思路：
// 1.链表转数组，数组首尾判断是否为回文串
func isPalindrome(head *ListNode) bool {
	// base case
	if head == nil || head.Next == nil {
		return true
	}
	// 链表中的元素转为int数组
	valList := make([]int, 0)
	for p := head; p != nil; {
		valList = append(valList, p.Val)
		p = p.Next
	}
	return check(valList)
}

// v2解法，递归
// 1.递归会先走到链表的最后节点，我们叫他right
// 2.递归外面设立一个全局变量left指针
// 3.双指针比较判断回文
func isPalindromeV2(head *ListNode) bool {
	left := head
	// 定义递归函数
	var dfs func(right *ListNode) bool
	dfs = func(right *ListNode) bool {
		// 递归结束条件
		if right == nil {
			return true
		}
		// 一直往后走
		if !dfs(right.Next) {
			return false
		}
		// 走到最后了，判断首尾是否相等
		if left.Val != right.Val {
			return false
		}
		// 首指针往后走
		left = left.Next
		return true
	}
	return dfs(head)
}

// check 数组是否为回文
func check(nums []int) bool {
	for i, j := 0, len(nums)-1; i < len(nums) && j >= 0; {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}
