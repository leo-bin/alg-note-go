package 其他类

import "testing"

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
