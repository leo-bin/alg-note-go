package 回溯递归

import (
	"reflect"
	"testing"
)

// 电话号码的字母组合，17，mid
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合
// 答案可以按 任意顺序 返回
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母
// 2【abc】，3【def】，4【ghi】，5【jkl】，6【mno】，7【pqrs】，8【tuv】，9【wxyz】
//
// 示例 1：
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 示例 2：
// 输入：digits = ""
// 输出：[]
//
// 示例 3：
// 输入：digits = "2"
// 输出：["a","b","c"]

func Test_LetterCombPhoneNumber(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Digits       string
		ExpectResult []string
	}{
		0: {
			Name:         "case 1",
			Digits:       "23",
			ExpectResult: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		2: {
			Name:         "case 2",
			Digits:       "",
			ExpectResult: []string{},
		},
		3: {
			Name:         "case 3",
			Digits:       "2",
			ExpectResult: []string{"a", "b", "c"},
		},
	}

	for i, testCase := range testCaseList {
		digits := testCase.Digits
		expect := testCase.ExpectResult
		actual := letterCombinations(digits)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

var phoneStrMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// letterCombinations 回溯
func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	var (
		res       = make([]string, 0)
		backTrace func(digits string, i int, curPath string)
	)
	backTrace = func(digits string, i int, curPath string) {
		if i == len(digits) {
			res = append(res, curPath)
			return
		}
		curPhone := digits[i]
		for _, str := range phoneStrMap[curPhone] {
			backTrace(digits, i+1, curPath+string(str))
		}
	}
	// 执行回溯
	backTrace(digits, 0, "")
	return res
}
