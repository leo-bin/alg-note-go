package 双指针解决数组问题

import (
	"reflect"
	"testing"
)

// 移除零值 leetcode-283 easy
// 链接：https://leetcode.cn/problems/move-zeroes/
// 题目介绍：
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
//
// 示例 2:
// 输入: nums = [0]
// 输出: [0]
func Test_RemoveZero(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		ExpectResult []int
	}{
		0: {
			Name:         "case1",
			Nums:         []int{0, 1, 0, 3, 12},
			ExpectResult: []int{1, 3, 12, 0, 0},
		},
		1: {
			Name:         "case2",
			Nums:         []int{0},
			ExpectResult: []int{0},
		},
	}
	for i, testCase := range testCaseList {
		nums := testCase.Nums
		expectResult := testCase.ExpectResult
		removeZero(nums)
		if !reflect.DeepEqual(nums, expectResult) {
			t.Errorf("test case %v,not passed,expectResult=%v||actualResult=%v", i, expectResult, nums)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// removeZero 快慢指针
// 慢指针记录可以替换的位置，快指针往前寻找可以替换的时机（当快指针指向的元素不等于0时）
func removeZero(nums []int) {
	// base case
	if len(nums) <= 0 {
		return
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return
}
