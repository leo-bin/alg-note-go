package 二叉树类问题

import (
	"math"
	"testing"
)

func Test_CountTreeNodes(t *testing.T) {

}

// countNodes dfs 一次遍历统计就完事，时间复杂度O（n）
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// countNodesV2 dfs+数学，时间复杂度O（logN*logN）
// 思路：
// 1.满二叉树的节点个数直接等于：2^h-1（其中h为树的深度）
// 2.也就是说我们优先判定当前树是不是满二叉树，如果是满二叉树，直接计算次方，如果不是，继续用dfs算
// 3.时间复杂度分析：一颗完全二叉树的子树，一定会有一颗树是满二叉树
func countNodesV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := root, root
	// 统计左子树的高度
	lH := 0
	for l != nil {
		l = l.Left
		lH++
	}
	// 统计右子树的高度
	rH := 0
	for r != nil {
		r = r.Right
		rH++
	}

	if lH == rH {
		return int(math.Pow(float64(2), float64(lH))) - 1
	}

	return countNodesV2(root.Left) + countNodesV2(root.Right) + 1
}
