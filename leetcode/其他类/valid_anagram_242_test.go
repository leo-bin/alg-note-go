package 其他类

import (
	"reflect"
	"testing"
)

// 有效的字母异位词，242，easy
// https://leetcode.cn/problems/valid-anagram/description/
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词
// 字母异位词是通过重新排列不同单词或短语的字母而形成的单词或短语，并使用所有原字母一次
//
// 示例 1:
// 输入: s = "anagram", t = "nagaram"
// 输出: true
//
// 示例 2:
// 输入: s = "rat", t = "car"
// 输出: false

func Test_isAnagram(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		T            string
		ExpectResult bool
	}{
		0: {
			Name:         "case 1",
			S:            "anagram",
			T:            "nagaram",
			ExpectResult: true,
		},
		1: {
			Name:         "case 1",
			S:            "rat",
			T:            "car",
			ExpectResult: false,
		},
	}

	for i, testCase := range testCaseList {
		ss, tt := testCase.S, testCase.T
		expect := testCase.ExpectResult
		actual := isAnagram(ss, tt)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// isAnagram hash
// 两个字符串a、b是否是字母异位词满足的条件有两个：
// 1、a和b的长度一致
// 2、a和b拥有一样的字符，且字符的出现次数一致
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make([]int, 26)
	// s累加一遍
	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}
	// t减回去
	for i := 0; i < len(t); i++ {
		record[t[i]-'a']--
		if record[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
