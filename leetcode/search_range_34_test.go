package leetcode

import (
	"testing"
)

// 在排序数组中查找元素的第一个和最后一个位置 leetcode 34
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
// 题目介绍
// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
// 示例 1：
//
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
// 示例 2：
//
// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]
// 示例 3：
//
// 输入：nums = [], target = 0
// 输出：[-1,-1]

func Test_SearchRange(t *testing.T) {
	testCaseList := []struct {
		Nums         []int
		Target       int
		ExpectResult []int
	}{
		0: {
			Nums:         []int{5, 7, 7, 8, 8, 10},
			Target:       8,
			ExpectResult: []int{3, 4},
		},
		1: {
			Nums:         []int{5, 7, 7, 8, 8, 10},
			Target:       6,
			ExpectResult: []int{-1, -1},
		},
		2: {
			Nums:         []int{},
			Target:       0,
			ExpectResult: []int{-1, -1},
		},
		3: {
			Nums:         []int{1},
			Target:       1,
			ExpectResult: []int{0, 0},
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		target := testCase.Target
		expectResult := testCase.ExpectResult
		actualResult := searchRange(nums, target)
		if len(actualResult) != 2 {
			t.Errorf("test case %v,failed,result num not 2", i)
		}
		if actualResult[0] == expectResult[0] && actualResult[1] == expectResult[1] {
			t.Logf("test case %v,passed", i)
		} else {
			t.Logf("test case %v,not passed", i)
		}
	}
}

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	// base case
	if len(nums) <= 0 {
		return result
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			// 继续往前找第一个位置
			i := mid - 1
			for i >= 0 && target == nums[i] {
				i--
			}
			j := mid + 1
			for j < len(nums) && target == nums[j] {
				j++
			}
			result[0], result[1] = i+1, j-1
			return result
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
