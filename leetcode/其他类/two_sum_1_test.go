package 其他类

import "testing"

// 两数之和，1，easy
// https://leetcode.cn/problems/two-sum/
// 给定一个整数数组 nums 和一个整数目标值 target
// 请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素
// 你可以按任意顺序返回答案
//
// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
//
// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
//
// 示例 3：
// 输入：nums = [3,3], target = 6
// 输出：[0,1]

func Test_TwoSum(t *testing.T) {

}

// twoSum 哈希表
func twoSum(nums []int, target int) []int {
	// base case
	if len(nums) <= 0 {
		return []int{-1, -1}
	}
	hashMap := make(map[int]int)
	for i, n := range nums {
		if x, ok := hashMap[target-n]; ok {
			return []int{i, x}
		}
		hashMap[n] = i
	}
	return []int{-1, -1}
}
