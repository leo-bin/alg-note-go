package 双指针解决数组问题

import "testing"

// 判断子序列,392,easy
// https://leetcode.cn/problems/is-subsequence/
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串
// 例如，"ace"是"abcde"的一个子序列，而"aec"不是
//
// 进阶：
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码
//
// 示例 1：
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
//
// 示例 2：
// 输入：s = "axc", t = "ahbgdc"
// 输出：false

func Test_IsSubSequence(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		T            string
		ExpectResult bool
	}{
		0: {"case 1", "abc", "ahbgdc", true},
		1: {"case 2", "axc", "ahbgdc", false},
	}

	for i, testCase := range testCaseList {
		s, tStr := testCase.S, testCase.T
		expect := testCase.ExpectResult
		actual := isSubSequence(s, tStr)
		if expect != actual {
			t.Errorf("test case %v,not passed,expect=%v,actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// isSubSequence 双指针
// 思路：
// 1.满足两个条件，s就是t的子序列
// 一）s的所有字符都在t中
// 二）s的所有的字符的顺序在t中的相对顺序保持一致，比如有s="abc"，t="aebicd"
// s字符的位置=[a（0）<b（1）<c（2）]，同样t字符的位置=[a（0）<b（2）<c（4）]
// 使用双指针技巧，分别从s和t的的首位开始往后匹配，当s[i]==t[j]时。说明匹配成功，同时往后移动
// 如果不相同，需要从t的后面开始匹配，也就是j往后移动
// 最后判断i是否走完了s
func isSubSequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] != t[j] {
			j++
		} else {
			i++
			j++
		}
	}
	return i == m
}
