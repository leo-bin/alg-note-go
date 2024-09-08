package 动态规划问题

import "testing"

// 最长回文子串 leetcode-5 mid
// 链接：https://leetcode.cn/problems/longest-palindromic-substring/description/
// 题目介绍：
// 给你一个字符串 s，找到 s 中最长的回文子串
// 回文子串：正着读和反着读都一样的字符串
// 比如说字符串 aba 和 abba 都是回文串，因为它们对称，反过来还是和本身一样
// 反之，字符串 abac 就不是回文串
//
// 示例1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案
//
// 示例2：
// 输入：s = "cbbd"
// 输出："bb"
func Test_LongestPalindrome(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		ExpectResult string
	}{
		0: {
			Name:         "case 1",
			S:            "babad",
			ExpectResult: "bab",
		},
		1: {
			Name:         "case 2",
			S:            "cbbd",
			ExpectResult: "bb",
		},
	}

	for i, testCase := range testCaseList {
		s := testCase.S
		expectResult := testCase.ExpectResult
		actualResult := longestPalindromeV2(s)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// longestPalindrome 动态规划
// dp思路：
func longestPalindrome(s string) string {
	return ""
}

// longestPalindromeV2 双指针
// 思路：
// 利用双指针技巧，枚举所有奇数长度和偶数长度的回文子串，记录最长的子串即可
func longestPalindromeV2(s string) string {
	var maxResult string
	for i := 0; i < len(s); i++ {
		// 1.从i开始，找奇数长度的最长回文子串
		s1 := palindrome(s, i, i)
		// 2.从i开始，找偶数长度的最长回文子串
		s2 := palindrome(s, i, i+1)
		// 3.记录最长子串
		if len(s1) > len(maxResult) {
			maxResult = s1
		}
		if len(s2) > len(maxResult) {
			maxResult = s2
		}
	}
	return maxResult
}

// palindrome 按照指定范围查找回文字串
func palindrome(s string, l int, r int) string {
	// 是回文串，继续往中间找
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// 此时的l-r就是指定l和r之间的最长回文串
	return s[l+1 : r]
}
