package 二叉树类问题

import "testing"

// 二叉树的最近公共祖先，leetcode-236，mid
// 链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
// 介绍：
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先
// 百度百科中最近公共祖先的定义为：
// 对于有根树 T 的两个节点 p、q
// 最近公共祖先表示为一个节点x,满足x是p、q祖先且x的深度尽可能大(一个节点也可是它自己祖先)

func Test_LowestCommonAncestor(t *testing.T) {

}

// lowestCommonAncestor dfs 查找
// 思路：
// 1.本质上就是在二叉树中进行递归查找p和q，如果找到了其中任意一个直接返回当前根节点
// 2.如如果从当前根节点出发找它的左右子树，如果左右都能查到结果，说明当前根节点就是最近的祖先
// 3.否则就返回找到的那个节点就行
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findAncestor(root, p.Val, q.Val)
}

// findAncestor
func findAncestor(root *TreeNode, p, q int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到p和q中的任意一个就行
	if root.Val == p || root.Val == q {
		return root
	}
	left := findAncestor(root.Left, p, q)
	right := findAncestor(root.Right, p, q)
	// 左右子树都能查找，说明当前节点就是祖先节点
	if left != nil && right != nil {
		return root
	}
	// 哪个不为空，返回哪个
	if left == nil {
		return right
	}
	return left
}
