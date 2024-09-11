package 二叉树类问题

import "testing"

// 二叉树的右侧视图，leetcode-199，mid
// 链接：https://leetcode.cn/problems/binary-tree-right-side-view/description/
// 题目介绍：
// 给定一个二叉树的 根节点 root
// 想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值
//
// 示例1：
// 输入: [1,2,3,null,5,null,4]
// 输出: [1,3,4]
//
// 示例2：
// 输入: [1,null,3]
// 输出: [1,3]

func Test_RightSideView(t *testing.T) {

}

// rightSideView dfs
// 思路：
// 每次优先遍历右节点而非左节点，实际上保证了“某一层第一次被访问的节点一定在这一层的最右边”
// 因此只要能判断某一层是不是被第一次访问就行，怎么判断？加一个depth参数就可以
func rightSideView(root *TreeNode) []int {
	var depth int
	var res []int
	traverseRightSideView(root, depth, &res)
	return res
}

func traverseRightSideView(root *TreeNode, depth int, res *[]int) {
	if root == nil {
		return
	}
	// 每一层最右边的元素只有一个，也就是res此时的元素个数一定等于遍历深度，说明此元素就是最右边的元素（从右往左遍历）
	if len(*res) == depth {
		*res = append(*res, root.Val)
	}
	traverseRightSideView(root.Right, depth+1, res)
	traverseRightSideView(root.Left, depth+1, res)
}
