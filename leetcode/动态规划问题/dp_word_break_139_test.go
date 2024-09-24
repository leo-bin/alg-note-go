package 动态规划问题

import "testing"

// 单次拆分，leetcode-139，mid
// 链接：https://leetcode.cn/problems/word-break/description
// 介绍：
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典
// 如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用
//
// 示例 1：
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
// 示例 2：
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//
// 示例 3：
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false

func Test_WordBreak(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		WordDict     []string
		ExpectResult bool
	}{
		0: {
			Name:         "case 1",
			S:            "leetcode",
			WordDict:     []string{"leet", "code"},
			ExpectResult: true,
		},
		1: {
			Name:         "case 2",
			S:            "applepenapple",
			WordDict:     []string{"apple", "pen"},
			ExpectResult: true,
		},
		2: {
			Name:         "case 3",
			S:            "catsandog",
			WordDict:     []string{"cats", "dog", "sand", "and", "cat"},
			ExpectResult: false,
		},
	}

	for i, testCase := range testCaseList {
		s := testCase.S
		wordDict := testCase.WordDict
		expectResult := testCase.ExpectResult
		actualResult := wordBreak(s, wordDict)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// wordBreak 回溯 会超时，时间复杂度O(2^N * M * N) 会超时
// 思路：
// 1.每一个单词都找一遍，
func wordBreak(s string, wordDict []string) bool {
	var res bool
	backTraceWordBreak(0, s, wordDict, &res)
	return res
}

func backTraceWordBreak(i int, s string, wordDict []string, found *bool) {
	// 找到了，直接返回
	if *found {
		return
	}
	if i == len(s) {
		*found = true
		return
	}
	for _, word := range wordDict {
		if i+len(word) <= len(s) && s[i:i+len(word)] == word {
			backTraceWordBreak(i+len(word), s, wordDict, found)
		}
	}
	return
}

// wordBreakDp dp
// 思路：
// 1.
func wordBreakDp(s string, wordDict []string) bool {
	return false
}
