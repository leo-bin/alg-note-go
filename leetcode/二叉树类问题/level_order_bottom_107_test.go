package 二叉树类问题

import "testing"

// 二叉树的层序遍历 ii，leetcode-107，mid
// 链接：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
// 给你二叉树的根节点 root ，返回其节点值自底向上的层序遍历
// 即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历
//
// 示例1:
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[15,7],[9,20],[3]]
//
// 示例2:
// 输入：root = [1]
// 输出：[[1]]
//
// 示例3:
// 输入：root = []
// 输出：[]

func Test_LevelOrderBottom(t *testing.T) {

}

// levelOrderBottom 层序遍历
func levelOrderBottom(root *TreeNode) [][]int {
	// base case
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		var levelRes []int
		levelSize := len(q)
		// 收集当前层的所有元素
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			levelRes = append(levelRes, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		//  将当前层的所有元素放入队头，保证倒序
		res = append([][]int{levelRes}, res...)
	}

	return res
}
