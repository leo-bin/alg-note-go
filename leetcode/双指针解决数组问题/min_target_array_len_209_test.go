package 双指针解决数组问题

import (
	"math"
	"reflect"
	"testing"
)

// 最小子数组，209，mid
// 给定一个含有 n 个正整数的数组和一个正整数 target
// 找出该数组中满足其总和大于等于 target 的长度最小的子数组
// [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/
//
// 示例 1：
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
// 示例 2：
// 输入：target = 4, nums = [1,4,4]
// 输出：1
//
// 示例 3：
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0

func Test_MinTargetArrayLen(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Target       int
		Nums         []int
		ExpectResult int
	}{
		0: {"case 1", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		1: {"case 2", 4, []int{1, 4, 4}, 1},
		2: {"case 3", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	for i, testCase := range testCaseList {
		target, nums := testCase.Target, testCase.Nums
		expect := testCase.ExpectResult
		actual := minTargetArrayLenV2(target, nums)
		if !reflect.DeepEqual(expect, actual) {
			t.Logf("test case %v,passed", i)
		} else {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		}
	}
}

// minTargetArrayLen 暴力 超时
func minTargetArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				minLen = minArrayLen(minLen, j-i+1)
			}
		}
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

// minTargetArrayLenV2 滑动窗口
func minTargetArrayLenV2(target int, nums []int) int {
	l, r, sum, minLen := 0, 0, 0, math.MaxInt
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			minLen = minArrayLen(minLen, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func minArrayLen(x, y int) int {
	if x > y {
		return y
	}
	return x
}
