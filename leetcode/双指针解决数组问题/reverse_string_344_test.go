package 双指针解决数组问题

import (
	"reflect"
	"testing"
)

// 反转字符串，leetcode-344，easy
// 链接：https://leetcode.cn/problems/reverse-string/description/
// 题目介绍：
// 编写一个函数，其作用是将输入的字符串反转过来
// 输入字符串以字符数组 s 的形式给出
// 不要给另外数组分配额外空间，你必须原地修改输入数组、用 O(1) 空间解决问题
//
// 示例1：
// 输入：s = ["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
//
// 示例 2：
// 输入：s = ["H","a","n","n","a","h"]
// 输出：["h","a","n","n","a","H"]

func Test_ReverseString(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            []byte
		ExpectResult []byte
	}{
		0: {
			Name:         "case 1",
			S:            []byte{'h', 'e', 'l', 'l', 'o'},
			ExpectResult: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		1: {
			Name:         "case 2",
			S:            []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			ExpectResult: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for i, testCase := range testCaseList {
		s := testCase.S
		expectResult := testCase.ExpectResult
		reverseString(s)
		actualResult := s
		if !reflect.DeepEqual(actualResult, expectResult) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// reverseString 反转字符串
// 思路：
// 双指针，两边交换，直到指针碰面
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
