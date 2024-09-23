package 动态规划问题

import "testing"

// 爬楼梯，leetcode-70，easy
// 链接：https://leetcode.cn/problems/climbing-stairs/description
// 介绍：
// 假设你正在爬楼梯，需要n阶你才能到达楼顶
// 每次你可以爬1或2个台阶
// 你有多少种不同的方法可以爬到楼顶呢？
// 示例 1：
// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶
// 2. 2 阶
//
// 示例 2：
// 输入：n = 3
// 输出：3
// 解释：有三种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶 + 1 阶
// 2. 1 阶 + 2 阶
// 3. 2 阶 + 1 阶
//
// 提示： 1 <= n <= 45

func Test_ClimbStairs(t *testing.T) {

}

// f(n)=f(n-1)+f(n-2)
func climbStairs(n int) int {
	// base case
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	// dp打表
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
