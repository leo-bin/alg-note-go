package 动态规划问题

import (
	"math"
	"testing"
)

// 买卖股票的最佳时机i，121，easy
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
// 给定一个数组prices，它的第i个元素 prices[i] 表示一支给定股票第 i 天的价格
// 你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票
// 设计一个算法来计算你所能获取的最大利润
// 返回你可以从这笔交易中获取的最大利润
// 如果你不能获取任何利润，返回 0
//
// 示例 1：
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5
// 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票
//
// 示例 2：
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

func Test_MaxProfit(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Prices       []int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Prices:       []int{7, 1, 5, 3, 6, 4},
			ExpectResult: 5,
		},
		1: {
			Name:         "case 2",
			Prices:       []int{7, 6, 4, 3, 1},
			ExpectResult: 0,
		},
	}

	for i, testCase := range testCaseList {
		prices := testCase.Prices
		expectResult := testCase.ExpectResult
		actualResult := maxProfitDp(prices)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// maxProfit 暴力解法，穷举所有的买卖情况，找到最大收益，时间复杂度O（接近n*n）
func maxProfit(prices []int) int {
	maxRes := math.MinInt
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			curProfit := prices[j] - prices[i]
			if curProfit <= 0 {
				continue
			}
			maxRes = maxP(maxRes, curProfit)
		}
	}
	if maxRes == math.MinInt {
		return 0
	}
	return maxRes
}

// maxProfitDp 有点类似贪心
// 思路：
// 1.只能买卖一次，求全局最大收益，那就得保证是在历史最低点买入，历史最高点卖出（按时间线性排序）
// 2.我们只需要遍历一遍数组，1.记录到当前为止的历史最低点，2.使用当前的价格-历史最低点得到当前为止的最大收益
// 3.遍历完，此时的最大收益就是我们要的结果
func maxProfitDp(prices []int) int {
	minPrice, maxRes := math.MaxInt, math.MinInt
	for i := 0; i < len(prices); i++ {
		minPrice = minP(minPrice, prices[i])
		maxRes = maxP(maxRes, prices[i]-minPrice)
	}
	if maxRes == math.MinInt {
		return 0
	}
	return maxRes
}

func minP(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxP(x, y int) int {
	if x > y {
		return x
	}
	return y
}
