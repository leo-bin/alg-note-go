package 二叉树类问题

import "math"

// isValidBST 验证二叉搜索树，98，mid
// https://leetcode.cn/problems/validate-binary-search-tree/description/
// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树
// 有效 二叉搜索树定义如下：
// 节点的左子树只包含 小于 当前节点的数
// 节点的右子树只包含 大于 当前节点的数
// 所有左子树和右子树自身必须也是二叉搜索树
// 思路：
// 1.递归，需要注意根节点的值需要传递给左右子树进行判定，本身是一个区间判断，不是简单的判断某个根节点的值是否满足就行
// 2.多传递2个参数，min和max，代表以当前节点为基准，最小的节点和最大的节点
func isValidBST(root *TreeNode) bool {
	// 递归结束条件
	if root == nil {
		return true
	}
	return checkBST(root, nil, nil)
}

func checkBST(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val < min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return checkBST(root.Left, min, root) && checkBST(root.Right, root, max)
}

// isValidBSTV2 v2 中序遍历
// 思路：
// 1.利用二叉搜索树的中序遍历是有序的特性
func isValidBSTV2(root *TreeNode) bool {
	prev := math.MinInt
	var (
		inorder func(node *TreeNode)
		flag    bool = true
	)
	inorder = func(node *TreeNode) {
		// 递归结束条件
		if node == nil {
			return
		}
		inorder(node.Left)
		if node.Val <= prev {
			flag = false
		}
		if !flag {
			return
		}
		prev = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return flag
}
