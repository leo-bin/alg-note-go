package 二叉树类问题

import (
	"strconv"
	"strings"
	"testing"
)

// 二叉树根节点到叶子节点数字之和，leetcode-129，mid
// 链接：https://leetcode.cn/problems/sum-root-to-leaf-numbers/
// 题目介绍：
// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个0到9之间的数字
// 每条从根节点到叶节点的路径都代表一个数字：
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123
// 计算从根节点到叶节点生成的 所有数字之和
// 叶节点 是指没有子节点的节点
//
// 示例1：
// 输入：root = [1,2,3]
// 输出：25
// 解释：
// 从根到叶子节点路径 1->2 代表数字 12
// 从根到叶子节点路径 1->3 代表数字 13
// 因此，数字总和 = 12 + 13 = 25
//
// 示例2：
// 输入：root = [4,9,0,5,1]
// 输出：1026
// 解释：
// 从根到叶子节点路径 4->9->5 代表数字 495
// 从根到叶子节点路径 4->9->1 代表数字 491
// 从根到叶子节点路径 4->0 代表数字 40
// 因此，数字总和 = 495 + 491 + 40 = 1026

func Test_SumNumbers(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Root         *TreeNode
		ExpectResult int
	}{
		0: {
			Name: "case 1",
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			ExpectResult: 25,
		},
		1: {
			Name: "case 2",
			Root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			ExpectResult: 1026,
		},
	}
	for i, testCase := range testCaseList {
		root := testCase.Root
		expectResult := testCase.ExpectResult
		actualResult := sumNumbers(root)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// sumNumbers dfs
// 思路：
// 整体是dfs遍历的框架
// 像这种要记录中间结果得到最终结果的，一般就弄2个变量，一个是临时变量tmp，用来记录遍历的子结果，一个变量res用来记录最终的结果
func sumNumbers(root *TreeNode) int {
	var path []string
	var res int
	traverseSumNumbers(root, &path, &res)
	return res
}

func traverseSumNumbers(root *TreeNode, path *[]string, res *int) {
	if root == nil {
		return
	}
	// 到达叶子节点
	if root.Left == nil && root.Right == nil {
		// 记录当前节点的值
		*path = append(*path, strconv.Itoa(root.Val))
		// 用当前路径的数字累加到res中
		pathVal, _ := strconv.Atoi(strings.Join(*path, ""))
		*res += pathVal
		// 清空当前节点的值
		*path = (*path)[:len(*path)-1]
		return
	}
	// 记录当前节点的值
	*path = append(*path, strconv.Itoa(root.Val))
	traverseSumNumbers(root.Left, path, res)
	traverseSumNumbers(root.Right, path, res)
	// 清空当前节点的值
	*path = (*path)[:len(*path)-1]
}
