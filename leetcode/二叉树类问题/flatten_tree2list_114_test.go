package 二叉树类问题

import "testing"

func Test_FlattenTree(t *testing.T) {

}

// flatten dfs 前序遍历+构建链表，时空复杂度都是O（n）
func flatten(root *TreeNode) {
	dummy := &TreeNode{Val: -1}
	p := dummy
	preBuild(root, p)
	root = dummy.Right
}

func preBuild(root, p *TreeNode) {
	if root == nil {
		return
	}
	p.Right = &TreeNode{Val: root.Val}
	p = p.Right

	preBuild(root.Left, p)
	preBuild(root.Right, p)
}

// flattenV2
func flattenV2(root *TreeNode) {

}
