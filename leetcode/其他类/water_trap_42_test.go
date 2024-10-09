package 其他类

import (
	"testing"
)

// 接雨水，42，hard
// https://leetcode.cn/problems/trapping-rain-water/description/
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图
// 计算按此排列的柱子，下雨之后能接多少雨水
//
// 示例1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可接6个单位雨水
//
// 示例 2：
// 输入：height = [4,2,0,3,2,5]
// 输出：9

func Test_WaterTrap(t *testing.T) {

}

// trap 暴力模拟，时间复杂度O（n）
// 思路：
// 1.每一个柱子能接的最大单位的雨水，取决于它的左右两边柱子的最大高度，假设分别为lmax和rmax
// 2.第i个柱子的所能接的最大单位的雨水=min(lmax,rmax)-h[i]，累计每一个柱子所能接住的最大单位雨水就是最终结果
// 3.注意不用考虑第一个柱子和最后一个柱子，因为这两个柱子左右两边没有遮挡，无法接住雨水
func trap(height []int) int {
	// base case
	n := len(height)
	if n <= 1 {
		return 0
	}
	// 1.提前计算好每一个位置的左右两边最高的柱子
	lMax, rMax := make([]int, n), make([]int, n)
	lMax[0] = height[0]
	rMax[n-1] = height[n-1]
	// 2.从左往右开始找每一个位置的左边的最大高度
	for i := 1; i < n; i++ {
		lMax[i] = maxTrap(lMax[i-1], height[i])
	}
	// 3.从右往左开始找每一个位置的右边的最大高度
	for i := n - 2; i >= 0; i-- {
		rMax[i] = maxTrap(rMax[i+1], height[i])
	}
	// 4.累计计算结果
	res := 0
	for i := 1; i < n; i++ {
		res += minTrap(lMax[i], rMax[i]) - height[i]
	}
	return res
}

func minTrap(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxTrap(x, y int) int {
	if x > y {
		return x
	}
	return y
}
