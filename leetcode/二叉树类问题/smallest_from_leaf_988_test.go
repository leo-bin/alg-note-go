package 二叉树类问题

import "testing"

// 从叶子结点开始的最小子字符串，leetcode-988，mid
// 链接：https://leetcode.cn/problems/smallest-string-starting-from-leaf/description/
//
// 题目介绍：
// 给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个 [0, 25] 范围内的值，分别代表字母 'a' 到 'z'
// 返回 按字典序最小 的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束
// 注：字符串中任何较短的前缀在 字典序上 都是 较小 的：
// 例如，在字典序上 "ab" 比 "aba" 要小。叶结点是指没有子结点的结点。
// 节点的叶节点是没有子节点的节点
//
// 示例1：
// 输入：root = [0,1,2,3,4,3,4]
// 输出："dba"
//
// 示例2：
// 输入：root = [25,1,3,1,3,0,2]
// 输出："adz"
//
// 示例3：
// 输入：root = [2,2,1,null,1,0,null,0]
// 输出："abc"

func Test_SmallestFromLeaf(t *testing.T) {

}

// smallestFromLead dfs
// 套用dfs框架就行
// 1.需要对val进行字符换算
// 2.字符是从叶子节点到根节点
// 3.在go中，string可直接比大小
func smallestFromLead(root *TreeNode) string {
	var res string
	var path []byte
	traverseSmallestFromLeaf(root, &path, &res)
	return res
}

func traverseSmallestFromLeaf(root *TreeNode, path *[]byte, res *string) {
	if root == nil {
		return
	}
	// 找到叶子结点，并将叶子节点作为最后一个值加入path
	if root.Left == nil && root.Right == nil {
		// 加入队头
		*path = append([]byte{byte('a' + root.Val)}, *path...)
		str := string(*path)
		// 更新最小str
		if *res == "" || str < *res {
			*res = str
		}

		// 清除path
		*path = (*path)[1:len(*path)]

		return
	}

	// 将当前节点值加入path中
	*path = append([]byte{byte('a' + root.Val)}, *path...)
	traverseSmallestFromLeaf(root.Left, path, res)
	traverseSmallestFromLeaf(root.Right, path, res)
	// 清除path
	*path = (*path)[1:len(*path)]
}
