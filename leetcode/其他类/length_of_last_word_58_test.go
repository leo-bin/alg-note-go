package 其他类

import (
	"testing"
)

// 最后一个单词的长度，58，easy
// https://leetcode.cn/problems/length-of-last-word/description/
// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开
// 返回字符串中 最后一个 单词的长度
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串
//
// 示例 1：
// 输入：s = "Hello World"
// 输出：5
// 解释：最后一个单词是“World”，长度为 5。
//
// 示例 2：
// 输入：s = "   fly me   to   the moon  "
// 输出：4
// 解释：最后一个单词是“moon”，长度为 4
//
// 示例 3：
// 输入：s = "luffy is still joyboy"
// 输出：6
// 解释：最后一个单词是长度为 6 的“joyboy”

func Test_LengthOfLastWord(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			S:            "Hello World",
			ExpectResult: 5,
		},
		1: {
			Name:         "case 2",
			S:            "   fly me   to   the moon  ",
			ExpectResult: 4,
		},
		2: {
			Name:         "case 3",
			S:            "luffy is still joyboy",
			ExpectResult: 6,
		},
	}

	for i, testCase := range testCaseList {
		s := testCase.S
		expectResult := testCase.ExpectResult
		actualResult := lengthOfLastWord(s)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// lengthOfLastWord 字符串匹配
// 思路：
// 1.遇到空格认为空格前面的字符组成一个单词，append进res数组中
// 2.返回res数组的最后一个元素即可
func lengthOfLastWord(s string) int {
	// base case
	if len(s) <= 0 {
		return 0
	}
	res := make([]string, 0)
	var curWord []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(curWord) > 0 {
				res = append(res, string(curWord))
				curWord = make([]byte, 0)
			}
		} else {
			curWord = append(curWord, s[i])
		}
	}
	// 最后一个元素
	if len(curWord) > 0 {
		res = append(res, string(curWord))
	}
	return len(res[len(res)-1])
}

// lengthOfLastWordV2 反向遍历
// 思路：
// 1.既然要求最后一个单词的长度，那就从后面往前面遍历，找到第一个单词返回即可
func lengthOfLastWordV2(s string) int {
	// 先把空格去除
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	res := 0
	for index >= 0 && s[index] != ' ' {
		res++
		index--
	}
	return res
}
