package 其他类

import (
	"strconv"
	"testing"
)

// 字符串相乘，43，mid
// https://leetcode.cn/problems/multiply-strings/description/
// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积
// 它们的乘积也表示为字符串形式
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数
//
// 示例 1:
// 输入: num1 = "2", num2 = "3"
// 输出: "6"
//
// 示例 2:
// 输入: num1 = "123", num2 = "456"
// 输出: "56088"

func Test_StringMulti(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Num1         string
		Num2         string
		ExpectResult string
	}{
		0: {
			Name:         "case 1",
			Num1:         "145",
			Num2:         "45",
			ExpectResult: "6525",
		},
		1: {
			Name:         "case 2",
			Num1:         "123",
			Num2:         "456",
			ExpectResult: "56088",
		},
	}

	for i, testCase := range testCaseList {
		num1, num2 := testCase.Num1, testCase.Num2
		expectResult := testCase.ExpectResult
		actualResult := multiply(num1, num2)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// multiply 双指针模拟
// 思路：
// 1.模拟竖式乘法计算过程，使用数组来存储临时结果
// 2.注意求进位用除数，求余位用余数
// 3.最后注意结果数组的前缀可能会有0，需要去掉
func multiply(num1, num2 string) string {
	if len(num1) <= 0 || len(num2) <= 0 {
		return ""
	}
	// 1.模拟竖式乘法的计算过程
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			pos1, pos2 := i+j, i+j+1
			// 当前的乘积要和上一位的进位相加
			sum := mul + res[pos2]
			res[pos1] += sum / 10
			res[pos2] = sum % 10
		}
	}
	// 2.去掉res所有的前缀0并转换为string
	i := 0
	for ; i < len(res); i++ {
		if res[i] != 0 {
			break
		}
	}
	// 3.res转成string
	var resStr string
	for ; i < len(res); i++ {
		resStr += strconv.Itoa(res[i])
	}
	if len(resStr) <= 0 {
		return "0"
	}

	return resStr
}
