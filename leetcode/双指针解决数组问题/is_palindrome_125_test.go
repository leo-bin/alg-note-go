package 双指针解决数组问题

import (
	"strings"
	"testing"
)

// 验证回文串，125，easy
// https://leetcode.cn/problems/valid-palindrome/description/
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样
// 则可以认为该短语是一个 回文串
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false
//
// 示例 1：
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
//
// 示例 2：
// 输入：s = "race a car"
// 输出：false
// 解释："raceacar" 不是回文串。
//
// 示例 3：
// 输入：s = " "
// 输出：true
// 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
// 由于空字符串正着反着读都一样，所以是回文串

func Test_IsPalindrome(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		ExpectResult bool
	}{
		0: {"case 1", "A man, a plan, a canal: Panama", true},
		1: {"case 2", "race a car", false},
		2: {"case 3", " ", true},
	}

	for i, testCase := range testCaseList {
		s, expec := testCase.S, testCase.ExpectResult
		actual := isPalindrome(s)
		if expec != actual {
			t.Errorf("test case %v,not passed,expec=%v,but actual=%v", i, expec, actual)
		} else {
			t.Logf("test cade %v,passed", i)
		}
	}
}

// isPalindrome 验证回文串 时空复杂度都是O（1）
// 思路：
// 1.双指针i和j分别从头和尾对s进行匹配，当s【i】!=s[j]时返回false不是回文串
// 2.当s全部判断完毕后，返回true
func isPalindrome(s string) bool {
	// 全转小写
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		// 找有效字符
		for i < len(s) && !isValidStr(s[i]) {
			i++
		}
		for j >= 0 && !isValidStr(s[j]) {
			j--
		}
		if i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}

// isValidStr 验证是否是有效字符
func isValidStr(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
