package 二分查找问题

import "testing"

// 搜索插入的位置，leetcode 35
// https://leetcode.cn/problems/search-insert-position/description/
// 题目描述：
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置
// 请必须使用时间复杂度为 O(log n) 的算法
//
// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
//
// 示例 2:
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
//
// 示例 3:
// 输入: nums = [1,3,5,6], target = 7
// 输出: 4

func Test_SearchInsert(t *testing.T) {
	testCaseList := []struct {
		Nums         []int
		Target       int
		ExpectResult int
	}{
		0: {
			Nums:         []int{1, 3, 5, 6},
			Target:       5,
			ExpectResult: 2,
		},
		1: {
			Nums:         []int{1, 3, 5, 6},
			Target:       2,
			ExpectResult: 1,
		},
		2: {
			Nums:         []int{1, 3, 5, 6},
			Target:       7,
			ExpectResult: 4,
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		target := testCase.Target
		expectResult := testCase.ExpectResult
		if searchInsert(nums, target) == expectResult {
			t.Logf("test case %v,passed", i)
		} else {
			t.Logf("test case %v,not passed", i)
		}
	}
}

func searchInsert(nums []int, target int) int {
	// base case
	if len(nums) <= 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
