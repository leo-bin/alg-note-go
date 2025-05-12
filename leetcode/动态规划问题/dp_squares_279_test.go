package 动态规划问题

// 完全平方数，279,mid
// https://leetcode.cn/problems/perfect-squares/description/
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量
// 完全平方数 是一个整数，其值等于另一个整数的平方
// 换句话说，其值等于一个整数自乘的积
// 例如，1、4、9 和 16 都是完全平方数
// 而 3 和 11 不是
func numSquares(n int) int {
	// base case
	if n == 1 {
		return 1
	}
	// 1.定义dp数组，表示数为i时少需要的平方数个数
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[j] = minInt(dp[j], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
