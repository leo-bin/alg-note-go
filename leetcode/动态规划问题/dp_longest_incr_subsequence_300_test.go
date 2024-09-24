package 动态规划问题

import (
	"math"
	"testing"
)

// 最长递增子序列，leetcode-300，mid
// 链接：https://leetcode.cn/problems/longest-increasing-subsequence/description/
// 介绍：
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列
//
// 示例 1：
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4
//
// 示例 2：
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
//
// 示例 3：
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1

func Test_LIS(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Nums:         []int{10, 9, 2, 5, 3, 7, 101, 18},
			ExpectResult: 4,
		},
		1: {
			Name:         "case 2",
			Nums:         []int{0, 1, 0, 3, 2, 3},
			ExpectResult: 4,
		},
		2: {
			Name:         "case 3",
			Nums:         []int{7, 7, 7, 7, 7, 7, 7},
			ExpectResult: 1,
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		expectResult := testCase.ExpectResult
		actualResult := lisDp(nums)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// lisDp dp
// 思路：
// 1.定义dp[i]为以nums[i]元素为尾部的最长子序列长度，遍历一遍nums数组，数组dp的最大值就是最长子序列的长度
// 2.如果nums[i]>nums[i-1]，那么dp[i]=max(dp[i],dp[i-1]+1)，后面一个元素比前一个元素要大，说明nums[i-1]和nums[i]本身就构成序列
func lisDp(nums []int) int {
	// 1.初始化dp表
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	// 2.固定一个元素，以当前元素开始从前往后找一遍
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 更新当前位置的最大值
				dp[i] = maxLisDp(dp[i], dp[j]+1)
			}
		}
	}

	// 3.找最大值并返回
	maxRes := math.MinInt
	for i := 0; i < len(dp); i++ {
		maxRes = maxLisDp(maxRes, dp[i])
	}

	return maxRes
}

func maxLisDp(x, y int) int {
	if x > y {
		return x
	}
	return y
}
