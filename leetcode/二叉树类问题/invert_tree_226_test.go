package 二叉树类问题

import "testing"

func Test_invertTree(t *testing.T) {

}

// invertTree dfs
// 思路：
// 直接dfs遍历，每次构造一个新的根节点，这个新根节点的左子树用旧的根节点的右子树构建的结果，右子树相反即可
func invertTree(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = invertTree(root.Right)
	newRoot.Right = invertTree(root.Left)

	return newRoot
}
