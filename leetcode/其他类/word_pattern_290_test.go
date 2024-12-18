package 其他类

import (
	"reflect"
	"strings"
	"testing"
)

// 单词规律，290，easy
// https://leetcode.cn/problems/word-pattern/
// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律
// 这里的 遵循 指完全匹配
// 例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律
//
// 示例1:
// 输入: pattern = "abba", s = "dog cat cat dog"
// 输出: true
//
// 示例 2:
// 输入:pattern = "abba", s = "dog cat cat fish"
// 输出: false
//
// 示例 3:
// 输入: pattern = "aaaa", s = "dog cat cat dog"
// 输出: false

func Test_WordPatten(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Pattern      string
		S            string
		ExpectResult bool
	}{
		0: {
			Name:         "case 1",
			Pattern:      "abba",
			S:            "dog cat cat dog",
			ExpectResult: true,
		},
		1: {
			Name:         "case 2",
			Pattern:      "abba",
			S:            "dog cat cat fish",
			ExpectResult: false,
		},
		2: {
			Name:         "case 3",
			Pattern:      "aaaa",
			S:            "dog cat cat dog",
			ExpectResult: false,
		},
	}
	for i, testCase := range testCaseList {
		p, s := testCase.Pattern, testCase.S
		expect := testCase.ExpectResult
		actual := wordPattern(p, s)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// wordPattern hash
func wordPattern(pattern string, s string) bool {
	// 1.格式化
	wordList := strings.Split(s, " ")
	if len(wordList) != len(pattern) {
		return false
	}
	// 2.build hash table
	pHashTable := make(map[byte]string)
	sHashTable := make(map[string]byte)
	// 3.match
	for i, w := range wordList {
		p := pattern[i]
		if sHashTable[w] > 0 && sHashTable[w] != p || len(pHashTable[p]) > 0 && pHashTable[p] != w {
			return false
		}
		pHashTable[p] = w
		sHashTable[w] = p
	}
	return true
}
