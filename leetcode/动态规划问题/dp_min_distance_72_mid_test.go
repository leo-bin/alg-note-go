package 动态规划问题

import "testing"

// 编辑距离，72，hard
// https://leetcode.cn/problems/edit-distance/description/
// 示例 1：
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
//
// 示例 2：
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')

func Test_Min(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S1           string
		S2           string
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			S1:           "horse",
			S2:           "ros",
			ExpectResult: 3,
		},
		1: {
			Name:         "case 2",
			S1:           "intention",
			S2:           "execution",
			ExpectResult: 5,
		},
	}
	for i, testCase := range testCaseList {
		s1 := testCase.S1
		s2 := testCase.S2
		expectResult := testCase.ExpectResult
		actualResult := minDistanceDfs(s1, s2)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// minDistanceDfs dfs+哈希表
func minDistanceDfs(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			memo[i][j] = -1
		}
	}
	i, j := len(word1)-1, len(word2)-1
	return dfsMinDistance(word1, i, word2, j, memo)
}

func dfsMinDistance(s1 string, i int, s2 string, j int, memo [][]int) int {
	// base case
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	// find return
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s1[i] == s2[j] {
		memo[i][j] = dfsMinDistance(s1, i-1, s2, j-1, memo)
	} else {
		memo[i][j] = min3(
			// insert
			dfsMinDistance(s1, i, s2, j-1, memo)+1,
			// del
			dfsMinDistance(s1, i-1, s2, j, memo)+1,
			// replace
			dfsMinDistance(s1, i-1, s2, j-1, memo)+1,
		)
	}
	return memo[i][j]
}

// minDistance dp
// 思路：
// 1.定义dp[i][j]为word1[0-i]和word2[0-j]的最小编辑距离
// 2.假设word1和word2的长度的分别为s1和s2，假设word1[0]和word2[0]为空字符串
// 3.那么就有dp[0][0-s2]=0,1,2,3...s2，dp[0-s1][0]=0,1,2,3...s1（画一个dp数组的表格就明白了）
// 4.我们从i=j=1开始，分别比较word1[i]和word2[j]中的所有字符，比较结果有两种：1.两个字符相同，2.两个字符不相同
// 第一种情况：dp[i][j]=dp[i-1][j-1]
// 第二种情况:考虑有3种编辑方法
// 3种编辑方法：
// 1.在word1[i]插入一个元素word[j]，此时word1[i]==word2[j]，此时的dp[i][j]=dp[i][j-1]+1，word2[j]被匹配，找word[j-1]的最小值
// 2.把word1[i]删除，此时word1[i]==word2[j]，此时dp[i][j]=dp[i-1][j]+1，找前一个，也就是word1[i-1]的最小值
// 3.把word1[i]替换成word2[j]，此时word1[i]==word2[j]，此时dp[i][j]=dp[i-1][j-1]+1
// 此时的dp[i][j]的最小值直接在3种编辑方法的结果取最小值就行
// 最后的dp[s1][s2]就是我们要最终结果
func minDistanceDp(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 1.定义dp数组，预留一位，假设word1[0]和word2[0]是空串
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	// 2.初始化dp数组
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}
	// 3.比较所有字符
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min3(
					// 在word1[i]插入一个word2[j]
					dp[i][j-1]+1,
					// 把word1[i]删了
					dp[i-1][j]+1,
					// 把word1[i]替换成word2[j]
					dp[i-1][j-1]+1,
				)
			}
		}
	}
	return dp[m][n]
}

// min3 return x,y,z smallest
func min3(x, y, z int) int {
	if x < y && x < z {
		return x
	}
	if y < x && y < z {
		return y
	}
	return z
}
