package 双指针解决数组问题

import (
	"math"
	"testing"
)

// 盛最多水的容器，11，mid
// https://leetcode.cn/problems/container-with-most-water/
// 给定一个长度为 n 的整数数组 height
// 有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水
// 返回容器可以储存的最大水量
//
// 示例1：
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]
// 在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49
//
// 示例2：
// 输入：height = [1,1]
// 输出：1

func Test_MaxArea(t *testing.T) {

}

// maxArea 双指针找最大值
// 思路：
// 1.双指针i和j从两边收缩判断面积最大
// 2.i和j此时的面积计算公式=min(h[i],h[j])*(j-i)
// 3.注意当找到max值时，需要往高度小的那边收缩，因为只有高度小的那一边才有可能往高度大的线靠近，从而是使得面积更大
func maxArea(height []int) int {
	// base case
	if len(height) <= 0 {
		return 0
	}
	l, r := 0, len(height)-1
	maxResult := math.MinInt
	for l < r {
		tmpResult := minIntArea(height[l], height[r]) * (r - l)
		maxResult = maxIntArea(maxResult, tmpResult)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxResult
}

func minIntArea(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxIntArea(x, y int) int {
	if x > y {
		return x
	}
	return y
}
