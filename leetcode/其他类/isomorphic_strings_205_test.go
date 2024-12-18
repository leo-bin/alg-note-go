package 其他类

import (
	"reflect"
	"testing"
)

// 同构字符串，205，easy
// https://leetcode.cn/problems/isomorphic-strings
// 给定两个字符串 s 和 t ，判断它们是否是同构的
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序
// 不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身
//
// 示例 1:
// 输入：s = "egg", t = "add"
// 输出：true
//
// 示例 2：
// 输入：s = "foo", t = "bar"
// 输出：false
//
// 示例 3：
// 输入：s = "paper", t = "title"
// 输出：true

func Test_isIsomorphic(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		T            string
		ExpectResult bool
	}{
		0: {
			Name:         "case 1",
			S:            "egg",
			T:            "add",
			ExpectResult: true,
		},
		1: {
			Name:         "case 2",
			S:            "foo",
			T:            "bar",
			ExpectResult: false,
		},
		2: {
			Name:         "case 3",
			S:            "paper",
			T:            "title",
			ExpectResult: true,
		},
		3: {
			Name:         "case 4",
			S:            "badc",
			T:            "baba",
			ExpectResult: false,
		},
		4: {
			Name:         "case 5",
			S:            "egcd",
			T:            "adfd",
			ExpectResult: false,
		},
	}

	for i, testCase := range testCaseList {
		ss, tt := testCase.S, testCase.T
		expect := testCase.ExpectResult
		actual := isIsomorphic(ss, tt)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test caee %v,passed", i)
		}
	}
}

// isIsomorphic hash
func isIsomorphic(s string, t string) bool {
	sHashTable := make(map[byte]byte, len(s))
	tHashTable := make(map[byte]byte, len(t))
	for i := 0; i < len(s); i++ {
		x, y := s[i], t[i]
		if sHashTable[x] > 0 && sHashTable[x] != y || tHashTable[y] > 0 && tHashTable[y] != x {
			return false
		}
		sHashTable[x] = y
		tHashTable[y] = x
	}
	return true
}
