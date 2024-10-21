package 双指针解决数组问题

import (
	"sort"
	"testing"
)

// 多数元素，169，easy
// https://leetcode.cn/problems/majority-element/description/
// 给定一个大小为 n 的数组 nums
// 返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素
//
// 示例 1：
// 输入：nums = [3,2,3]
// 输出：3
//
// 示例 2：
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2

func Test_MajorityElement(t *testing.T) {

}

// majorityElement map
// 思路：
func majorityElement(nums []int) int {
	target := len(nums) / 2
	recordMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		record, ok := recordMap[nums[i]]
		if !ok {
			record = 1
			recordMap[nums[i]] = record
		} else {
			record++
			recordMap[nums[i]] = record
		}
		if record > target {
			return nums[i]
		}
	}
	return -1
}

// majorityElement map
// 思路：
// 快慢指针，寻找重复元素出现个数，如果个数大于多数，直接返回
// 2,2,1,1,1,2,2 -> [1,1,1,2,2,2,2]
func majorityElementV2(nums []int) int {
	target := len(nums) / 2
	sort.Ints(nums)
	i, j := 0, 0
	record := 0
	for j < len(nums) {
		if nums[i] != nums[j] {
			if record > target {
				return nums[i]
			} else {
				record = 0
				i = j
			}
		}
		record++
		j++
	}
	return nums[len(nums)-1]
}
