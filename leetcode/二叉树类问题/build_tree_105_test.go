package 二叉树类问题

import "testing"

// 从前序和中序遍历中构造二叉树，leetcode-105，mid
// 链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
// 题目介绍：
// 给定两个整数数组 preorder 和 inorder
// 其中preorder是二叉树的先序遍历,inorder是同一棵树的中序遍历
// 请构造二叉树并返回其根节点
//
// 示例1：
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//
// 示例2：
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]

func Test_buildTree(t *testing.T) {

}

// buildTree dfs+前序和中序遍历的特点
// 思路：
// 1.前序遍历一定总是先遍历根节点，然后在遍历左右子树
// 2.中序遍历一定是总是先遍历根节点的左子树，然后遍历根节点和根节点的右子树
// 3.所以，前序遍历顺序中，第一个元素一定是根节点
// 4.根据前序遍历找到根节点，然后在中序遍历顺序中找到根节点的左右子节点的个数
// 5.通过左右子树的数量得到左右子节点的元素范围，构建左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	// 构造根节点
	root := &TreeNode{Val: preorder[0]}
	i := 0
	// 在中序遍历顺序中找到根节点的位置
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 根据跟节点所在位置计算根节点左右子树的元素格式
	leftSize := len(inorder[:i])
	// 构建左右子树
	root.Left = buildTree(preorder[1:leftSize+1], inorder[:i])
	root.Right = buildTree(preorder[leftSize+1:], inorder[i+1:])

	return root
}
