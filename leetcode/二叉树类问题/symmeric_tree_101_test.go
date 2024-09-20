package 二叉树类问题

import (
	"reflect"
	"testing"
)

// 对称二叉树，leetcode-102，easy
// 链接：https://leetcode.cn/problems/symmetric-tree/description/
// 介绍：
// 给你一个二叉树的根节点root,检查它是否轴对称
//
// 示例1：
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
// 示例2：
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false

func Test_SymmetricTree(t *testing.T) {

}

// isSymmetric dfs解法
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

// check dfs 判定两棵树是否对称
func check(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// isSymmetricV2 bfs层序遍历
// 思路：
// 1.如果是对称的，那么从左到右的层序结果和从右到左的层序结果是一致的
// 2.需要考虑一些特殊情况，比如左节点是空，但右节点有值（反之依然），这种情况要进行补齐，人工构造一个虚拟节点加入到下一层中（需要在完全二叉树的情况下才能满足）
func isSymmetricV2(root *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		var levelRes []int
		var levelResReverse []int
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			levelRes = append(levelRes, cur.Val)
			levelResReverse = append([]int{cur.Val}, levelResReverse...)

			if cur.Val == 0 {
				continue
			}

			if cur.Left == nil {
				q = append(q, &TreeNode{Val: 0})
			} else {
				q = append(q, cur.Left)
			}

			if cur.Right == nil {
				q = append(q, &TreeNode{Val: 0})
			} else {
				q = append(q, cur.Right)
			}
		}
		// 如果是对称的，那么从左到右的层序结果和从右到左的层序结果是一致的
		if !reflect.DeepEqual(levelRes, levelResReverse) {
			return false
		}
	}
	return true
}
