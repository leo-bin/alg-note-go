package 二叉树类问题

import "testing"

// 二叉树最长连续序列，leetcode-298，mid
// 链接（会员）：https://leetcode.cn/problems/binary-tree-longest-consecutive-sequence/
//
// 题目介绍：
// 给你一棵指定的二叉树的根节点root
// 请你计算其中最长连续序列路径的长度
// 最长连续序列路径 是依次递增1的路径
// 该路径，可以是从某个初始节点到树中任意节点，通过「父-子」关系连接而产生的任意路径
// 且必须从父节点到子节点，反过来是不可以的
//
// 示例1
// 输入：root = [1,null,3,2,4,null,null,null,5]
// 输出：3
// 解释：当中，最长连续序列是 3-4-5 ，所以返回结果为 3
//
// 示例2
// 输入：root = [2,null,3,2,null,1]
// 输出：2
// 解释：当中，最长连续序列是 2-3
// 注意，不是 3-2-1，所以返回 2
func Test_LongestConsecutive(t *testing.T) {

}

// longestConsecutive dfs
// 思路：
// 依旧是套用dfs框架，我们需要3个临时变量，分别用来记录上一个节点的值、当前有效的连续序列长度、最长序列大小
func longestConsecutive(root *TreeNode) int {
	var tmpLen int = 0
	var res int = 0
	traverseLongestConsecutive(root, 0, tmpLen, &res)
	return res
}

func traverseLongestConsecutive(root *TreeNode, parentVal int, tmpLen int, res *int) {
	if root == nil {
		return
	}
	// 判断当前节点和上一个节点是不是构成连续
	if (root.Val - 1) == parentVal {
		// 记录构成连续的个数
		tmpLen++
	} else {
		// 不连续的话，从当前节点开始，算1个
		tmpLen = 1
	}
	// 更新最大序列值
	if tmpLen > *res {
		*res = tmpLen
	}
	traverseLongestConsecutive(root.Left, root.Val, tmpLen, res)
	traverseLongestConsecutive(root.Right, root.Val, tmpLen, res)
}
