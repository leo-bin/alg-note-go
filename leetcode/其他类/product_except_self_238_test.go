package 其他类

import "testing"

// 除自身之外的乘积和，238，mid
// https://leetcode.cn/problems/product-of-array-except-self/
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题
//
// 示例 1:
// 输入: nums = [1,2,3,4]  234、134、124、123
// 输出: [24,12,8,6]
//
// 示例 2:
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]

func Test_ProductExceptSelf(t *testing.T) {

}

// productExceptSelf
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1
	// 计算左边的累计乘积
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	// 计算右边的累计乘积
	right[len(nums)-1] = 1
	for j := len(nums) - 2; j >= 0; j-- {
		right[j] = nums[j+1] * right[j+1]
	}
	// 左右两边乘起来
	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i]
	}
	return res
}
