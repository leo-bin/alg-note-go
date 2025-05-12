package 二叉树类问题

import "testing"

// 二叉搜索树中第K小的元素，230，mid
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
// 给定一个二叉搜索树的根节点 root ，和一个整数 k
// 请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）

func TestKthSmallest(t *testing.T) {

}

// 思路：
// 1.二叉搜索树的中序遍历结果是天然升序的
// 2.遍历到中间节点的时候记录下rank就行
func kthSmallest(root *TreeNode, k int) int {
	var (
		dfs  func(node *TreeNode)
		res  int
		rank int
	)
	dfs = func(node *TreeNode) {
		// 递归结束条件
		if node == nil {
			return
		}
		dfs(node.Left)
		rank++
		// 匹配到k了，直接返回
		if rank == k {
			res = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}
