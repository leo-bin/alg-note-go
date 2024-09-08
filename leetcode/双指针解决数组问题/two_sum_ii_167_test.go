package 双指针解决数组问题

import (
	"reflect"
	"testing"
)

// 两数之和，leetcode-167 mid
// 链接：https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
//
// 题目介绍：
// 给你一个下标从 1 开始的整数数组 numbers
// 该数组已按 非递减顺序排列,请你从数组中找出满足相加之和等于目标数 target 的两个数
// 如果设这两个数分别是numbers[index1]和numbers[index2],1 <= index1 < index2 <= numbers.length
// 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素
// 你所设计的解决方案必须只使用常量级的额外空间
//
// 示例1：
// 输入：numbers = [2,7,11,15], target = 9
// 输出：[1,2]
//
// 示例 2：
// 输入：numbers = [2,3,4], target = 6
// 输出：[1,3]
// 解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3]
//
// 示例 3：
// 输入：numbers = [-1,0], target = -1
// 输出：[1,2]
// 解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2]
func Test_TwoSumII(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Numbers      []int
		Target       int
		ExpectResult []int
	}{
		0: {
			Name:         "case 1",
			Numbers:      []int{2, 7, 11, 15},
			Target:       9,
			ExpectResult: []int{1, 2},
		},
		1: {
			Name:         "case 2",
			Numbers:      []int{2, 3, 4},
			Target:       6,
			ExpectResult: []int{1, 3},
		},
		2: {
			Name:         "case 3",
			Numbers:      []int{-1, 0},
			Target:       -1,
			ExpectResult: []int{1, 2},
		},
		3: {
			Name:         "case 4",
			Numbers:      []int{1, 2, 3, 4, 5},
			Target:       10,
			ExpectResult: []int{-1, -1},
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Numbers
		target := testCase.Target
		expectResult := testCase.ExpectResult
		actualResult := twoSum(nums, target)
		if !reflect.DeepEqual(actualResult, expectResult) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// 两数之和ii 双指针移动解决
// 思路：
// 类似于二分查找，左右两个指针同时往中间查找
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}
