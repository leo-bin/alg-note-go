package 双指针解决数组问题

import "testing"

// 删除有序数组中的重复项 leetcode-26 easy
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
// 题目介绍：
// 给你一个非严格递增排列的数组nums
// 请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要
// 返回 k
//
// 示例1：
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素
//
// 示例 2：
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素

func Test_RemoveDuplicates(t *testing.T) {

}

// removeDuplicates 快慢指针
// 思想：
// 1.数组有序，说明重复的元素一定是连续出现
// 2.要求原地修改，不能使用额外容器进行去重判断
// 3.快慢指针思想，快指针在前面探路，慢指针在后面进行修改和标记
// 两个指针从第一个元素开始一起前进，当快指针所在的元素和slow指针不一样，说明找到了下一段重复元素
// 此时slow指针往前走一步，同时用fast指针的元素和slow进行替换
func removeDuplicates(nums []int) int {
	// base case
	if len(nums) <= 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		// 下一个不重复的元素出现
		if nums[fast] != nums[slow] {
			slow++
			// 原地修改数组
			nums[slow] = nums[fast]
		}
		fast++
	}
	// 此时的0-slow就是修改后的数组
	return slow + 1
}
