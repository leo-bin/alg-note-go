package 二叉树类问题

import (
	"math"
	"testing"
)

// 构造最大二叉树，leetcode-654，mind
// 链接：https://leetcode.cn/problems/maximum-binary-tree/description/
// 题目介绍：
// 给定一个不重复的整数数组 nums
// 最大二叉树 可以用下面的算法从 nums 递归地构建:
// 创建一个根节点，其值为 nums 中的最大值
// 递归地在最大值 左边 的 子数组前缀上 构建左子树
// 递归地在最大值 右边 的 子数组后缀上 构建右子树
// 返回 nums 构建的 最大二叉树
//
// 示例1：
// 输入：nums = [3,2,1,6,0,5]
// 输出：[6,3,5,null,2,0,null,null,1]
//
// 示例2：
// 输入：nums = [3,2,1]
// 输出：[3,null,2,null,1]

func Test_ConstructMaxBinaryTree(t *testing.T) {

}

// constructMaxBinaryTree dfs
// 思路：
// 每次找数组中的最大值作为当前的根节点
// 分别遍历最大值左边的子数组和右边的子数组作为当前根节点的左右子树
func constructMaxBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	// 找当前数组的最大值
	maxValue := math.MinInt32
	maxValueIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			maxValueIndex = i
		}
	}
	root := &TreeNode{Val: maxValue}
	// 找当前根节点的左子树
	root.Left = constructMaxBinaryTree(nums[:maxValueIndex])
	root.Right = constructMaxBinaryTree(nums[maxValueIndex+1:])

	return root
}
