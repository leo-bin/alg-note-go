package 其他类

import (
	"reflect"
	"testing"
)

// 快乐数，202，easy
// https://leetcode.cn/problems/happy-number/description/
// 编写一个算法来判断一个数 n 是不是快乐数
// 「快乐数」 定义为：
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1
// 如果这个过程 结果为 1，那么这个数就是快乐数
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false
//
// 示例 1：
// 输入：n = 19
// 输出：true
// 解释：
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1
//
// 示例 2：
// 输入：n = 2
// 输出：false
func Test_IsHappy(t *testing.T) {
	testCaseList := []struct {
		Name         string
		N            int
		ExpectResult bool
	}{
		0: {
			Name:         "case 1",
			N:            19,
			ExpectResult: true,
		},
		1: {
			Name:         "case 2",
			N:            2,
			ExpectResult: false,
		},
	}

	for i, testCase := range testCaseList {
		n := testCase.N
		expect := testCase.ExpectResult
		actual := isHappy(n)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// isHappy hash
// 思路：
// 1、可以找下规律
// 2、假设n=19，流程是： 19 -> 82 -> 68 -> 100 ->1，最后找到n=1的时候所有数都不重复
// 3、假设n=2，流程是：2 -> 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4（重复）
func isHappy(n int) bool {
	hMap := make(map[int]struct{})
	for n != 1 {
		if _, ok := hMap[n]; !ok {
			hMap[n] = struct{}{}
		} else {
			return false
		}
		n = singleHappy(n)
	}
	return true
}

// singleHappy 计算单次happy数
func singleHappy(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
