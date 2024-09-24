package 动态规划问题

import "testing"

// 打家劫舍，leetcode-198，mid
// 介绍：
// 你是一个专业的小偷，计划偷窃沿街的房屋
// 每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警
// 给定一个代表每个房屋存放金额的非负整数数组
// 计算你不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额
//
// 示例 1：
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
// 偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
// 输入：[2,7,9,3,1]
// 输出：12
// 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
// 偷窃到的最高金额 = 2 + 9 + 1 = 12

func Test_Rob(t *testing.T) {

}

// rob dp
// 思路：
// 1.定义dp[i]表示第i间房能够抢到的最大金额
// 2.每次抢，只能有两种方案
// 一：抢第i间房，但是不能抢i-1的房间，此时能够抢到的最大金额就是dp[i-2]+nums[i]
// 二：不抢第i间房，那肯定抢i-1间房，此时的最大金额就是dp[i-1]
func rob(nums []int) int {
	// base case
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxRob(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = maxRob(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[len(nums)-1]
}

func maxRob(x, y int) int {
	if x > y {
		return x
	}
	return y
}
