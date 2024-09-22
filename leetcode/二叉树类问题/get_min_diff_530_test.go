package 二叉树类问题

import (
	"math"
	"testing"
)

// 二叉搜索树的最小绝对差值，leetcode-530，easy
// 链接：https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description
// 介绍：
// 给你一个二叉搜索树的根节点 root
// 返回树中任意两不同节点值之间的最小差值
// 差值是一个正数，其数值等于两值之差的绝对值

func Test_GetMinDiff(t *testing.T) {

}

// getMinDiff dfs
// 思路：
// 1.二叉搜索树的特点：中序遍历的顺序是递增的
// 2.所以求两个节点的最小绝对差值，本质上只需要在中序遍历顺序中找就行
func getMinDiff(root *TreeNode) int {
	min, pre := math.MaxInt, -1
	dfsMinDiff(root, &min, &pre)
	return min
}

func dfsMinDiff(root *TreeNode, min, pre *int) {
	// base case
	if root == nil {
		return
	}
	dfsMinDiff(root.Left, min, pre)

	if *pre != -1 {
		// 更新最小值
		curMin := int(math.Abs(float64(root.Val - *pre)))
		if curMin < *min {
			*min = curMin
		}
	}
	*pre = root.Val

	dfsMinDiff(root.Right, min, pre)
}
