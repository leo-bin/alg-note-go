package 二叉树类问题

import "testing"

// 从根到叶的二进制之和，leetcode-1022，easy
// 链接：https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/description/
// 题目介绍：
// 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
// 例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
// 对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
// 返回这些数字之和。题目数据保证答案是一个 32 位 整数
//
// 示例1:
// 输入：root = [1,0,1,0,1,0,1]
// 输出：22
// 解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
//
// 示例2:
// 输入：root = [0]
// 输出：0

func Test_SumRootToLeaf(t *testing.T) {

}

//type TreeNode struct {
// Val   int
// Left  *TreeNode
// Right *TreeNode
//}

// sumRootToLeaf dfs
// 思路：
// dfs遍历框架，这里需要注意下如何对多个二进制进行拼接和相加的操作
// 首先按照顺序，需要不断往做移，让当前节点值加入到path中，然后通过或操作进行相加
func sumRootToLeaf(root *TreeNode) int {
	var path int
	var res int
	traverseSumRootToLeaf(root, &path, &res)
	return res
}

func traverseSumRootToLeaf(root *TreeNode, path *int, res *int) {
	if root == nil {
		return
	}
	// 找到叶子节点
	if root.Left == nil && root.Right == nil {
		// 计算当前二进制路径的十进制的值，更新全局结果
		*res += *path<<1 | root.Val
		return
	}
	// 将当前节点加入到二进制路径里去
	*path = *path<<1 | root.Val
	traverseSumRootToLeaf(root.Left, path, res)
	traverseSumRootToLeaf(root.Right, path, res)
	// 撤销当前节点
	*path = *path >> 1
}
