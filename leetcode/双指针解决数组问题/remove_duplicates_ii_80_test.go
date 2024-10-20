package 双指针解决数组问题

import "testing"

// 删除有序数组中的重复项ii，80，mid
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素
// 使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成
//
// 示例 1：
// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
//
// 示例 2：
// 输入：nums = [0,0,1,1,1,1,2,3,3]
// 输出：7, nums = [0,0,1,1,2,3,3]
// 解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素

func Test_RemoveDuplicatesII(t *testing.T) {

}

// removeDuplicatesII 双指针
// 思路：
// 1.xx
func removeDuplicatesII(nums []int) int {
	// base case
	if len(nums) <= 0 {
		return 0
	}
	slow, fast := 0, 0
	count := 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2 {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
		count++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			count = 0
		}
	}
	return slow + 1
}
