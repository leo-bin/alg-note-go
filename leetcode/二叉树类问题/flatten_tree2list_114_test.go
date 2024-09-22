package 二叉树类问题

import "testing"

// 二叉树展开为链表，leetcode-114，mid
// 链接：
// 介绍：
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode
// 其中 right 子指针指向链表中下一个结点，而左子指针始终为 null
// 展开后的单链表应该与二叉树 先序遍历 顺序相同

func Test_FlattenTree(t *testing.T) {
	testCaseList := []struct {
		Name string
		Root *TreeNode
	}{
		0: {
			Name: "case 1",
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, testCase := range testCaseList {
		root := testCase.Root
		flattenV2(root)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	var prePath []*TreeNode
	preBuildV2(root, &prePath)
	for i := 1; i < len(prePath); i++ {
		pre, cur := prePath[i-1], prePath[i]
		// 构建链表关系
		pre.Left, pre.Right = nil, cur
	}
}

func preBuildV2(root *TreeNode, prePath *[]*TreeNode) {
	if root == nil {
		return
	}
	*prePath = append(*prePath, root)
	preBuildV2(root.Left, prePath)
	preBuildV2(root.Right, prePath)
}
