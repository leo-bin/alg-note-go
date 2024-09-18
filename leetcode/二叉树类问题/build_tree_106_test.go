package 二叉树类问题

import "testing"

// 从后序和中序遍历顺序中构建二叉树，leetcode-106，mid
// 链接：https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
// 题目介绍：
// 给定两个整数数组 inorder 和 postorder
// 其中inorder是二叉树的中序遍历，postorder是同一棵树的后序遍历
// 请你构造并返回这颗 二叉树
//
// 示例1：
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
//
// 示例2：
// 输入：inorder = [-1], postorder = [-1]
// 输出：[-1]

func Test_BuildTree106(t *testing.T) {

}

// buildTree106 dfs+后序和中序遍历的特点
// 思路：
// 1.后序遍历总是先遍历左右子树，最后遍历根节点，所以后续遍历顺序的最后一个元素一定是二叉树的根节点
// 2.中序遍历总是先遍历左子树，然后遍历根节点和右子树
// 3.先根据1的特点在后续遍历中找到根节点，然后在中序遍历中找到根节点的位置
// 4.有了根节点的位置就可推断出根节点左右子树的元素个数，进而构建左右子树
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(postorder) <= 0 {
		return nil
	}
	// 后序遍历的最后一个元素就是根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	// 根据根节点的位置确定根节点左子树的元素个数
	leftSize := len(inorder[:i])
	root.Left = buildTree106(inorder[:i], postorder[:leftSize])
	root.Right = buildTree106(inorder[i+1:], postorder[leftSize:len(postorder)-1])

	return root
}
