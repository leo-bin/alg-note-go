package 二叉树类问题

// 二叉树的最大直径，543，easy
// https://leetcode.cn/problems/diameter-of-binary-tree/
// 给你一棵二叉树的根节点，返回该树的直径
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度
// 这条路径可能经过也可能不经过根节点 root
// 两节点之间路径的 长度 由它们之间边数表示

// 思路：
// 1.两个节点之间的路径长度等价于求两个节点路径中经过了多少个节点-1
// 2.求两个节点路径中经过了多少个节点等价于求以某个根节点root开始，root的left有多少个节点+root的right有多少个节点+1（算上根节点之际）
func diameterOfBinaryTree(root *TreeNode) int {
	var (
		maxRes int
		dfs    func(node *TreeNode) int
	)
	dfs = func(node *TreeNode) int {
		// 递归结束条件
		if node == nil {
			return 0
		}
		lMax := dfs(node.Left)
		rMax := dfs(node.Right)
		maxRes = max(maxRes, lMax+rMax+1)
		return max(lMax, rMax) + 1
	}
	dfs(root)
	return maxRes - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
