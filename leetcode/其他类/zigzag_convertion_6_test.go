package 其他类

import (
	"reflect"
	"testing"
)

// Z字形变换，6，mid
// https://leetcode.cn/problems/zigzag-conversion/description/
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
// P   A   H   N
// A P L S I I G
// Y   I   R
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串
// 比如："PAHNAPLSIIGYIR"
// 请你实现这个将字符串进行指定行数变换的函数：
//
// 示例 1：
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
//
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
// 示例 3：
// 输入：s = "A", numRows = 1
// 输出："A"

func Test_ZigZagConv(t *testing.T) {
	testCaseList := []struct {
		Name         string
		S            string
		NumRows      int
		ExpectResult string
	}{
		0: {"case 1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		1: {"case 2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		2: {"case 3", "A", 1, "A"},
	}

	for i, testCase := range testCaseList {
		s, n := testCase.S, testCase.NumRows
		expectResult := testCase.ExpectResult
		actualResult := zigZagConvert(s, n)
		if !reflect.DeepEqual(expectResult, actualResult) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// zigZagConvert 模拟z字形走位构建每一层的元素
// 0 1 2 3 2 1 0 1 2 3
func zigZagConvert(s string, numRows int) string {
	// base case
	if numRows < 2 {
		return s
	}
	// 1.构造每一层的z列表
	zList := make([][]byte, numRows)
	pos, flag := 0, -1
	for i := 0; i < len(s); i++ {
		zList[pos] = append(zList[pos], s[i])
		// 重置方向
		if pos == numRows-1 || pos == 0 {
			flag = -flag
		}
		pos += flag
	}
	// 2.每一层的z按顺序打平构造结果返回
	res := ""
	for i := 0; i < len(zList); i++ {
		res += string(zList[i])
	}
	return res
}
