package 其他类

import (
	"strconv"
	"testing"
)

// 字符串相加，415，easy
// https://leetcode.cn/problems/add-strings/description/
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）
// 也不能直接将输入的字符串转换为整数形式。
//
// 示例 1：
// 输入：num1 = "11", num2 = "123"
// 输出："134"
//
// 示例 2：
// 输入：num1 = "456", num2 = "77"
// 输出："533"
//
// 示例 3：
// 输入：num1 = "0", num2 = "0"
// 输出："0"

func Test_StringAdd(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Num1         string
		Num2         string
		ExpectResult string
	}{
		0: {
			Name:         "case 1",
			Num1:         "11",
			Num2:         "123",
			ExpectResult: "134",
		},
		1: {
			Name:         "case 2",
			Num1:         "456",
			Num2:         "77",
			ExpectResult: "533",
		},
		2: {
			Name:         "case 3",
			Num1:         "0",
			Num2:         "0",
			ExpectResult: "0",
		},
	}

	for i, testCase := range testCaseList {
		num1, num2 := testCase.Num1, testCase.Num2
		expectResult := testCase.ExpectResult
		actualResult := addStrings(num1, num2)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// addStrings 双指针模拟
// 思路：
// 1.直接模拟相机的过程即可
func addStrings(num1, num2 string) string {
	x, y := len(num1), len(num2)
	// base case
	if x <= 0 || y <= 0 {
		return ""
	}
	add := 0
	resString := ""
	for i, j := x-1, y-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		a, b := 0, 0
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}
		result := a + b + add

		add = result / 10
		resString = strconv.Itoa(result%10) + resString
	}
	return resString
}
