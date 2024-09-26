package 动态规划问题

import "testing"

// 最长公共子序列，1143，mid
// https://leetcode.cn/problems/longest-common-subsequence/description/
// 介绍：
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度
// 如果不存在 公共子序列 ，返回 0
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序情况下删除某些字符（也可不删除任何字符）后组成的新字符串
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列
//
// 示例 1：
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
//
// 示例 2：
// 输入：text1 = "abc", text2 = "abc"
// 输出：3
// 解释：最长公共子序列是 "abc" ，它的长度为 3
//
// 示例 3：
// 输入：text1 = "abc", text2 = "def"
// 输出：0
// 解释：两个字符串没有公共子序列，返回 0

func Test_LCS(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Text1        string
		Text2        string
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Text1:        "abcde",
			Text2:        "ace",
			ExpectResult: 3,
		},
		1: {
			Name:         "case 2",
			Text1:        "abc",
			Text2:        "abc",
			ExpectResult: 3,
		},
		2: {
			Name:         "case 3",
			Text1:        "abc",
			Text2:        "def",
			ExpectResult: 0,
		},
	}

	for i, testCase := range testCaseList {
		text1 := testCase.Text1
		text2 := testCase.Text2
		expectResult := testCase.ExpectResult
		actualResult := longestCommonSubsequence(text1, text2)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// longestCommonSubsequence lcs dp解
// 思路：
// 1.定义dp[i][j]为text1[0-i]和text2[0-j]的最长公共子序列的长度
// 2.假设字符串的真实索引从1开始，也就是另text1[0]=text2[0]=""，text1和text2的长度分别为m和n
// 3.dp[0][0-j]=dp[i][0]=0，空串和任意一个有效字符比，最长公共子序列都是空串，也就是长度为0
// 4.我们从i=j=1开始遍历两个字符（有效走字符开始），会遇到两种情况：
// 第一种情况： text1[i] = text2[j]
// 公共子序列长度+1，也就是dp[i][j]=dp[i-1][j-1]+1
// 第二种情况：text1[i] != text2[j]
// 我们可以选择删除text1[i]，也可选择删除text2[j]，删除的意思就是使用i-1时的结果，也就是还没有text1[i]时的最长公共子序列的长度，text2[j]同理
// 此时的dp[i][j]=max(dp[i-1][j],dp[i][j-1])
// 遍历完后，最终，dp[m][n]就是最长公共子序列的长度
func longestCommonSubsequence(text1 string, text2 string) int {
	// 1.定义dp数组
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	// 2.初始化dp数组(可以不用，因为int数组初始化的默认值就是0)
	for i := 0; i < m+1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = 0
	}
	// 3.遍历两个字符
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxLCS(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// 4.返回结果
	return dp[m][n]
}

func maxLCS(x, y int) int {
	if x > y {
		return x
	}
	return y
}
