package 二叉树类问题

import "testing"

// 相同的树，leetcode-100，easy
// 链接：https://leetcode.cn/problems/same-tree/description/
// 介绍：
// 给你两棵二叉树的根节点 p 和 q
// 编写一个函数来检验这两棵树是否相同
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的
//
// 示例1：
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true
//
// 示例2：
// 输入：p = [1,2], q = [1,null,2]
// 输出：false

func Test_IsSameTree(t *testing.T) {

}

// isSameTree dfs
// 思路：
// dfs两棵树同时遍历，满足左右子树都是相同的条件即可
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// base case
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
