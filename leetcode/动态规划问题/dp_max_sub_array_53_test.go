package 动态规划问题

import (
	"math"
	"testing"
)

// 最大子数组和，53，mid
// https://leetcode.cn/problems/maximum-subarray/description/
// 介绍：
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素）
// 返回其最大和
// 子数组是数组中的一个连续部分。
//
// 示例 1：
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
// 示例 2：
// 输入：nums = [1]
// 输出：1
//
// 示例 3：
// 输入：nums = [5,4,-1,7,8]
// 输出：23

func Test_MaxSubArray(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Nums:         []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			ExpectResult: 6,
		},
		1: {
			Name:         "case 2",
			Nums:         []int{5, 4, -1, 7, 8},
			ExpectResult: 23,
		},
		2: {
			Name:         "case 3",
			Nums:         []int{1},
			ExpectResult: 1,
		},
	}
	for i, testCase := range testCaseList {
		nums := testCase.Nums
		expectResult := testCase.ExpectResult
		actualResult := maxSubArray(nums)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = nums[i]
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}
	// 返回最大值
	maxRes := math.MinInt
	for i := 0; i < len(dp); i++ {
		maxRes = max(maxRes, dp[i])
	}
	return maxRes
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
