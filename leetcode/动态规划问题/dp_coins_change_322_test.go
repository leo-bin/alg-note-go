package 动态规划问题

import (
	"math"
	"testing"
)

// 零钱兑换，leetcode-322，mid
// 链接：https://leetcode.cn/problems/coin-change/description/
// 介绍：
// 给你一个整数数组 coins ，表示不同面额的硬币
// 以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数
// 如果没有任何一种硬币组合能组成总金额，返回 -1
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
//
// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1
//
// 示例 3：
// 输入：coins = [1], amount = 0
// 输出：0

func Test_CoinsChange(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Coins        []int
		Amount       int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Coins:        []int{2},
			Amount:       3,
			ExpectResult: -1,
		},
		1: {
			Name:         "case 2",
			Coins:        []int{1, 2, 5},
			Amount:       11,
			ExpectResult: 3,
		},
	}

	for i, testCase := range testCaseList {
		coins := testCase.Coins
		amount := testCase.Amount
		expectResult := testCase.ExpectResult
		actualResult := coinsChangeDfs(coins, amount)
		if expectResult != actualResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// dfs+哈希表剪枝
func coinsChangeDfs(coins []int, amount int) int {
	resSet := make(map[int]int)
	return dfsCoinsChange(&resSet, coins, amount)
}

// dfsCoinsChange
func dfsCoinsChange(resSet *map[int]int, coins []int, amount int) int {
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if res, okRes := (*resSet)[amount]; okRes {
		return res
	}

	// 递归求解子问题的最优解
	minRes := math.MaxInt
	for _, curCoin := range coins {
		curRes := dfsCoinsChange(resSet, coins, amount-curCoin)
		if curRes < 0 {
			continue
		}
		minRes = min(minRes, curRes+1)
	}

	if minRes == math.MaxInt {
		minRes = -1
	}

	(*resSet)[amount] = minRes

	return minRes
}

// dp打表
func coinsChangeDp(coins []int, amount int) int {
	// 初始化dp表，金额为dp[i]表示当i=1、2、3、4的最小硬币数量
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		// 不能用maxInt，否则后面进行子问题求解的时候+1会导致整形溢出
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			// 子问题无解
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
