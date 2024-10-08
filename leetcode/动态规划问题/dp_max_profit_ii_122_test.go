package 动态规划问题

import "testing"

// 买卖股票的最佳时机，122，mid
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票
// 你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售
// 返回 你能获得的最大利润
//
// 示例 1：
// 输入：prices = [7,1,5,3,6,4]
// 输出：7
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
// 最大总利润为 4 + 3 = 7 。
//
// 示例 2：
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 最大总利润为 4 。
//
// 示例 3：
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。

func Test_MaxProfit2(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Prices       []int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Prices:       []int{7, 1, 5, 3, 6, 4},
			ExpectResult: 7,
		},
	}
	for i, testCase := range testCaseList {
		prices := testCase.Prices
		expectResult := testCase.ExpectResult
		actualResult := maxProfit2(prices)
		if expectResult != actualResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// maxProfit2 dp
// 思路：
// 1.定义dp[i][j]，i为能够交易的天数，j为当天是否持有股票，0=不持有，1=持有
// 2.dp[i][0]=当天不持有股票时的最大收益，dp[i][1]为当天持有股票时的最大收益，最后一天不持有股票的结果就是最大收益
// 3.每次我们可买入卖出，也就是有两种选择：
// 卖出时最大收益，也就是dp[i][0]=max(dp[i-1][0],dp[i-1][0]+prices[i])
// 买入时最大收益，也就是dp[i][1]=max(dp[i-1][1],dp[i-1][1]-prices[i])
func maxProfit2(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(prices); i++ {
		// base case
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// 当天卖掉
		dp[i][0] = maxP2(dp[i-1][0], dp[i-1][1]+prices[i])
		// 当天买入
		dp[i][1] = maxP2(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func maxP2(x, y int) int {
	if x > y {
		return x
	}
	return y
}
