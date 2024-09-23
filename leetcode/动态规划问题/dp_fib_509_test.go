package 动态规划问题

import (
	"testing"
)

// 斐波那契数列，leetcode-509，easy
// 链接：https://leetcode.cn/problems/fibonacci-number/description/
// 介绍：
// 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列
// 该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
// F(0) = 0，F(1) = 1
// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
// 给定 n ，请计算 F(n)
//
// 示例1:
// 输入：n = 2
// 输出：1
// 解释：F(2) = F(1) + F(0) = 1 + 0 = 1
//
// 示例 2：
// 输入：n = 3
// 输出：2
// 解释：F(3) = F(2) + F(1) = 1 + 1 = 2
//
// 示例 3：
// 输入：n = 4
// 输出：3
// 解释：F(4) = F(3) + F(2) = 2 + 1 = 3

func Test_Fib(t *testing.T) {
	testCaseList := []struct {
		Name         string
		N            int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			N:            3,
			ExpectResult: 2,
		},
		1: {
			Name:         "case 2",
			N:            2,
			ExpectResult: 1,
		},
		2: {
			Name:         "case 3",
			N:            4,
			ExpectResult: 3,
		},
	}

	for i, testCase := range testCaseList {
		n := testCase.N
		expectResult := testCase.ExpectResult
		//actualResult := fibDfs(n)
		//actualResult := fibDp(n)
		actualResult := fibDpV2(n)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// dfs+哈希表剪支
func fibDfs(n int) int {
	resSet := make(map[int]int)
	return dfsFib(&resSet, n)
}

func dfsFib(resSet *map[int]int, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 哈希表中有直接返回
	if res, okRes := (*resSet)[n]; okRes {
		return res
	}

	// 计算当前结果
	curRes := dfsFib(resSet, n-1) + dfsFib(resSet, n-2)

	// 保存记录
	(*resSet)[n] = curRes

	return curRes
}

// dp打表
func fibDp(n int) int {
	// base case
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// 初始化dp表
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	// 状态转移
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// dp+空间优化，减少不必要的空间存储
func fibDpV2(n int) int {
	// base case
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// 使用2个变量记录n-1和n-2的值
	f1, f2 := 1, 0
	var res int
	for i := 2; i <= n; i++ {
		res = f1 + f2
		f2 = f1
		f1 = res
	}

	return res
}
