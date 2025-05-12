package 二叉树类问题

import (
	"encoding/json"
	"fmt"
	"testing"
)

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

func Test_buildTree105(t *testing.T) {
	testCaseList := []struct {
		Name     string
		PreOrder []int
		inOrder  []int
	}{
		0: {
			Name:     "case 1",
			PreOrder: []int{3, 9, 20, 15, 7},
			inOrder:  []int{9, 3, 15, 20, 7},
		},
	}

	for _, testCase := range testCaseList {
		preorder := testCase.PreOrder
		inorder := testCase.inOrder
		root := buildTree105(preorder, inorder)
		result, _ := json.Marshal(root)
		fmt.Printf("root=%v", string(result))
	}
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// buildTree dfs+前序和中序遍历的特点
// 思路：
// 1.前序遍历一定总是先遍历根节点，然后在遍历左右子树
// 2.中序遍历一定是总是先遍历根节点的左子树，然后遍历根节点和根节点的右子树
// 3.所以，前序遍历顺序中，第一个元素一定是根节点
// 4.根据前序遍历找到根节点，然后在中序遍历顺序中找到根节点的左右子节点的个数
// 5.通过左右子树的数量得到左右子节点的元素范围，构建左右子树
func buildTree105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	// 前序遍历顺序的的第一个元素就是根节点
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 根据根节点所在的位置计算根节点左右子树的元素个数
	leftSize := len(inorder[:i])
	// 确定根节点左右元素的范围
	root.Left = buildTree105(preorder[1:leftSize+1], inorder[:i])
	root.Right = buildTree105(preorder[leftSize+1:], inorder[i+1:])

	return root
}

// 从先序和中序顺序构建二叉树，105,mid
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
// dfs
// 思路：
// 1.先序遍历的第一个元素就是根节点
// 2.中序遍历优先遍历左子树，在遍历根节点，然后是右子树
// 3.先通过先序遍历找到根节点
// 4.然后再中序遍历中找到根节点的左右子树数量
// 5.递归构建树
func buildTree(preorder []int, inorder []int) *TreeNode {
	// base case
	if len(preorder) <= 0 {
		return nil
	}
	// 构建中序遍历元素映射表
	hashInOrder := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		hashInOrder[inorder[i]] = i
	}
	var dfs func(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode
	dfs = func(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
		// 递归结束条件
		if preStart > preEnd {
			return nil
		}
		// 找根节点
		root := &TreeNode{Val: preorder[preStart]}
		rootIndex := hashInOrder[root.Val]
		leftSize := rootIndex - inStart

		// 分别构造左右子树
		root.Left = dfs(preorder, preStart+1, preStart+leftSize, inorder, inStart, rootIndex-1)
		root.Right = dfs(preorder, preStart+leftSize+1, preEnd, inorder, rootIndex+1, inEnd)

		return root
	}
	return dfs(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
