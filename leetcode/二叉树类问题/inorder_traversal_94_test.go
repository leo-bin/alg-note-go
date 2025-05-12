package 二叉树类问题

// 二叉树中序遍历，94，easy
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历
func inorderTraversal(root *TreeNode) []int {
	// base case
	if root == nil {
		return []int{}
	}
	var (
		traversal func(node *TreeNode)
		res       []int
	)
	// 定义递归函数
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历，先遍历left，在遍历根，在遍历right
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
