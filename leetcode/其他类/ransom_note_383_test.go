package 其他类

import (
	"reflect"
	"testing"
)

// 赎金信，383，easy
// https://leetcode.cn/problems/ransom-note/description/
// 给你两个字符串：ransomNote 和 magazine
// 判断 ransomNote 能不能由 magazine 里面的字符构成
// 如果可以，返回 true ；否则返回 false
// magazine 中的每个字符只能在 ransomNote 中使用一次
//
// 示例 1：
// 输入：ransomNote = "a", magazine = "b"
// 输出：false
//
// 示例 2：
// 输入：ransomNote = "aa", magazine = "ab"
// 输出：false
//
// 示例 3：
// 输入：ransomNote = "aa", magazine = "aab"
// 输出：true

func Test_CanConstruct(t *testing.T) {
	testCaseList := []struct {
		Name         string
		RansomNote   string
		Magazine     string
		ExpectResult bool
	}{
		0: {
			Name:         "case 1",
			RansomNote:   "a",
			Magazine:     "b",
			ExpectResult: false,
		},
		1: {
			Name:         "case 2",
			RansomNote:   "aa",
			Magazine:     "ab",
			ExpectResult: false,
		},
		2: {
			Name:         "case 3",
			RansomNote:   "aa",
			Magazine:     "aab",
			ExpectResult: true,
		},
	}
	for i, testCase := range testCaseList {
		ransomNote, magazine := testCase.RansomNote, testCase.Magazine
		expect := testCase.ExpectResult
		t.Run(testCase.Name, func(t *testing.T) {
			actual := canConstruct(ransomNote, magazine)
			if !reflect.DeepEqual(expect, actual) {
				t.Errorf("test case %v,not passed,exptect=%v,but actual=%v", i, expect, actual)
			} else {
				t.Logf("test case %v,passed", i)
			}
		})
	}
}

// canConstruct hash映射
func canConstruct(ransomNote string, magazine string) bool {
	// 1、build hash table
	strRecord := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		strRecord[magazine[i]-'a']++
	}
	// 2.judge
	for i := 0; i < len(ransomNote); i++ {
		if strRecord[ransomNote[i]-'a'] <= 0 {
			return false
		} else {
			strRecord[ransomNote[i]-'a']--
		}
	}
	return true
}
