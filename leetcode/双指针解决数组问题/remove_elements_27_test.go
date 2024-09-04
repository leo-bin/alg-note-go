package 双指针解决数组问题

import "testing"

// 移除元素 leetcode-27
// 链接：https://leetcode.cn/problems/remove-element/description/
// 题目介绍：
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。
// 元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量
// 假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
// 更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
// 返回 k
//
// 示例1:
// 输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2,_,_]
// 解释：你的函数函数应该返回 k = 2, 并且 nums 中的前两个元素均为 2。
// 你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。
//
// 示例 2：
// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
// 输出：5, nums = [0,1,4,0,3,_,_,_]
// 解释：你的函数应该返回 k = 5，并且 nums 中的前五个元素为 0,0,1,3,4。
// 注意这五个元素可以任意顺序返回。
// 你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）

func Test_RemoveElements(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		Val          int
		ExpectResult int
	}{
		0: {
			Name:         "case1",
			Nums:         []int{3, 2, 2, 3},
			Val:          3,
			ExpectResult: 2,
		},
		1: {
			Name:         "case1",
			Nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val:          2,
			ExpectResult: 5,
		},
	}
	for i, testCase := range testCaseList {
		nums := testCase.Nums
		val := testCase.Val
		expectResult := testCase.ExpectResult
		actualResult := removeElements(nums, val)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v||actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// removeElements 快慢指针
// 思路：
// 慢指针用来记录可替换的位置，快指针查询可替换的时机（不等于目标值即可进行替换）
func removeElements(nums []int, val int) int {
	// base case
	if len(nums) <= 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// 此时slow记录的就是不等于目标值的元素个数
	return slow
}
