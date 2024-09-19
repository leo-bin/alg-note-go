package 二叉树类问题

import "testing"

// 二叉树的锯齿形层序遍历，leetcode-103，mid
// 链接：https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
// 介绍：
// 给你二叉树的根节点 root ，返回其节点值的锯齿形层序遍历
// 即先从左往右，再从右往左进行下一层遍历
// 以此类推，层与层之间交替进行
//
// 示例1:
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
//
// 示例 2：
// 输入：root = [1]
// 输出：[[1]]
//
// 示例 3：
// 输入：root = []
// 输出：[]

func Test_ZigzagLevelOrder(t *testing.T) {

}

// zigzagLevelOrder z字形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	// base case
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := []*TreeNode{root}
	// 使用flag来控制是从左还是从右收集每一层结果
	flag := true
	for len(q) > 0 {
		var levelRes []int
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			if flag {
				// 从左往右
				levelRes = append(levelRes, cur.Val)
			} else {
				// 从右往左
				levelRes = append([]int{cur.Val}, levelRes...)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		// 反转标记
		flag = !flag
		// 收集每一层的所有元素
		res = append(res, levelRes)
	}

	return res
}
