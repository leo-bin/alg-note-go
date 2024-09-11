package 二叉树类问题

import "testing"

// 二叉树的最大深度，leetcode-104，easy
// 链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/description
// 题目介绍：
// 给定一个二叉树 root ，返回其最大深度
// 二叉树的最大深度是指从根节点到最远叶子节点的最长路径上的节点数
//
// 示例1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
//
// 示例2:
// 输入：root = [1,null,2]
// 输出：2

func Test_MaxDepth(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Root         *TreeNode
		ExpectResult int
	}{
		0: {
			Name: "case 1",
			Root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			ExpectResult: 3,
		},
	}

	for i, testCase := range testCaseList {
		root := testCase.Root
		expectResult := testCase.ExpectResult
		actualResult := maxDepth(root)
		if expectResult != actualResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth dfs回溯
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)

	return maxInt(leftMax, rightMax) + 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
