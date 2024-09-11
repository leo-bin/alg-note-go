package 二叉树类问题

import (
	"strconv"
	"strings"
	"testing"
)

// 二叉树的所有路径，leetcode-257，easy
// 链接：https://leetcode.cn/problems/binary-tree-paths/
// 题目介绍：
// 给你一个二叉树的根节点 root
// 按任意顺序 ，返回所有从根节点到叶子节点的路径
// 叶子节点 是指没有子节点的节点
//
// 示例1：
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
//
// 示例2：
// 输入：root = [1]
// 输出：["1"]
func Test_BinaryTreePaths(t *testing.T) {

}

func binaryTreePaths(root *TreeNode) []string {
	var path []string
	var res []string
	traverse(root, &path, &res)
	return res
}

// traverse dfs遍历一遍即可
func traverse(root *TreeNode, path *[]string, res *[]string) {
	if root == nil {
		return
	}
	// 左右子树都是空，说明是叶子节点，记录结果返回，清空path，继续后面的遍历
	if root.Left == nil && root.Right == nil {
		*path = append(*path, strconv.Itoa(root.Val))
		*res = append(*res, strings.Join(*path, "->"))
		*path = (*path)[:len(*path)-1]
		return
	}
	// 加入当前节点的值
	*path = append(*path, strconv.Itoa(root.Val))
	traverse(root.Left, path, res)
	traverse(root.Right, path, res)
	// 移除当前节点的值
	*path = (*path)[:len(*path)-1]
}
