package sort_and_search

import "testing"

// leetcode 704 ：二分查找
// https://leetcode.cn/problems/binary-search/
//
// 题目介绍：
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1
//
// 示例 1:
//
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//示例 2:
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1

func Test_BinarySearch(t *testing.T) {
	testCaseList := []struct {
		Nums         []int
		Target       int
		ExpectResult int
	}{
		0: {
			Nums:         []int{0, 1, 2, 3, 4},
			Target:       1,
			ExpectResult: 1,
		},
		1: {
			Nums:         []int{},
			Target:       2,
			ExpectResult: -1,
		},
		2: {
			Nums:         []int{0, 1, 2, 3, 4, 5},
			Target:       5,
			ExpectResult: 5,
		},
		3: {
			Nums:         []int{0, 1, 2, 3, 4, 5},
			Target:       6,
			ExpectResult: -1,
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		target := testCase.Target
		expectResult := testCase.ExpectResult
		if binarySearch(nums, target) == expectResult {
			t.Logf("test case %v, passed", i)
		} else {
			t.Logf("test case %v,failed", i)
		}
	}
}

// binarySearch 二分查找
func binarySearch(nums []int, target int) int {
	// base case
	if len(nums) <= 0 {
		return -1
	}

	leftIndex, rightIndex := 0, len(nums)-1
	for leftIndex <= rightIndex {
		midIndex := leftIndex + (rightIndex-leftIndex)/2
		if target == nums[midIndex] {
			return midIndex
		} else if target > nums[midIndex] {
			leftIndex = midIndex + 1
		} else {
			rightIndex = midIndex - 1
		}
	}

	return -1
}

// 个人总结：
// 1.核心思想就是利用数组有序，在一个范围内不断逼近最终的目标值，从中间那个值开始找
// 2.注意int和int相加可能会导致溢出（十进制在计算机中是以二进制存储，第一位是符号位，0表示正数，-1表示负数，超出二进制范围的，符号位会变成1，也就是负数）
// 3.
