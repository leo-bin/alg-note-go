package 其他类

import (
	"reflect"
	"strings"
	"testing"
)

// 反转字符串中的单词，151,mid
// https://leetcode.cn/problems/reverse-words-in-a-string/
// 给你一个字符串 s ，请你反转字符串中单词的顺序
//
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格
// 返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格
//
// 示例 1：
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
//
// 示例 2：
// 输入：s = "  hello world  "
// 输出："world hello"
// 解释：反转后的字符串中不能存在前导空格和尾随空格。
//
// 示例 3：
// 输入：s = "a good   example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。

func Test_ReverseWords(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		ExpectResult string
	}{
		0: {
			"case 1", "the sky is blue", "blue is sky the",
		},
		1: {
			"case 2", "  hello world  ", "world hello",
		},
		2: {
			"case 3", "a good   example", "example good a",
		},
	}
	for i, testCase := range testCaseList {
		s := testCase.S
		expectResult := testCase.ExpectResult
		actualResult := reverseWords(s)
		if !reflect.DeepEqual(expectResult, actualResult) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// reverseWord 额外数组存储反转后的word
func reverseWords(s string) string {
	// base case
	if len(s) <= 0 {
		return ""
	}

	var tmpWord []byte
	var res []string

	// 去掉前导空格
	i := 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	// 找完整的word，按照先入后出的方式存入res中
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			if len(tmpWord) > 0 {
				res = append([]string{string(tmpWord)}, res...)
				tmpWord = make([]byte, 0)
			}
			continue
		} else {
			tmpWord = append(tmpWord, s[i])
		}
	}
	if len(tmpWord) > 0 {
		res = append([]string{string(tmpWord)}, res...)
	}
	// 按空格分隔转string返回
	return strings.Join(res, " ")
}

// reverseWordsV2 原地反转
func reverseWordsV2(s string) string {
	return ""
}
