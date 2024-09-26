package 动态规划问题

import "testing"

// 两个字符串的最小删除距离，583，mid
// https://leetcode.cn/problems/delete-operation-for-two-strings/description/
// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数
// 每步可以删除任意一个字符串中的一个字符。
//
// 示例 1：
// 输入: word1 = "sea", word2 = "eat"
// 输出: 2
// 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
//
// 示例  2:
// 输入：word1 = "leetcode", word2 = "etco"
// 输出：4

func Test_MinDistance2(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S1           string
		S2           string
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			S1:           "sea",
			S2:           "eat",
			ExpectResult: 2,
		},
	}
	for i, testCase := range testCaseList {
		s1 := testCase.S1
		s2 := testCase.S2
		expectResult := testCase.ExpectResult
		actualResult := minDelDistance(s1, s2)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// minDelDistance 最小删除距离 dp
// 思路：
// 1.先求lcs，最小删除步数理论上只需要把两个字符剩下不重叠的元素都删除即可
func minDelDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	maxLcsLen := lcs(word1, word2)
	return (m - maxLcsLen) + (n - maxLcsLen)
}

func lcs(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxLcs(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func maxLcs(x, y int) int {
	if x > y {
		return x
	}
	return y
}
