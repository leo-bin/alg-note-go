package 其他类

import "testing"

// 跳跃游戏i，55，mid
// https://leetcode.cn/problems/jump-game/description/
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标
// 数组中的每个元素代表你在该位置可以跳跃的最大长度
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false
//
// 示例 1：
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标
//
// 示例 2：
// 输入：nums = [3,2,1,0,4]
// 输出：false
// 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标

func Test_JumpGame(t *testing.T) {

}

// jumpGame 贪心（类似于动态规划）
// 思路：
// 1.把所有位置的最远跳跃距离维护起来，最后判断最远距离是否大于等于最后一个位置来判断是否能跳到最后
// 2.中途如果出现最远距离小于当前的位置，说明跳不下去了
func jumpGame(nums []int) bool {
	maxJump := 0
	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}
		maxJump = maxJumpGame(maxJump, i+nums[i])
	}
	return maxJump >= len(nums)-1
}

func maxJumpGame(x, y int) int {
	if x > y {
		return x
	}
	return y
}
