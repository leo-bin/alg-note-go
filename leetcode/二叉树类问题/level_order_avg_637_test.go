package 二叉树类问题

import (
	"reflect"
	"testing"
)

// 二叉树的层平均值，leetcode-637，easy
// 链接：https://leetcode.cn/problems/average-of-levels-in-binary-tree/description/
// 介绍：
// 给定一个非空二叉树的根节点 root
// 以数组的形式返回每一层节点的平均值
// 与实际答案相差 10-5 以内的答案可以被接受
//
// 示例1:
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[3.00000,14.50000,11.00000]
// 解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11
// 因此返回 [3, 14.5, 11]
//
// 示例2:
// 输入：root = [3,9,20,15,7]
// 输出：[3.00000,14.50000,11.00000]

func Test_LevelOrderAvg(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Root         *TreeNode
		ExpectResult []float64
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
			ExpectResult: []float64{3.00000, 14.50000, 11.00000},
		},
	}
	for i, testCase := range testCaseList {
		root := testCase.Root
		expectResult := testCase.ExpectResult
		actualResult := avgLevels(root)
		if !reflect.DeepEqual(actualResult, expectResult) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,paassed", i)
		}
	}
}

func avgLevels(root *TreeNode) []float64 {
	// base case
	if root == nil {
		return []float64{}
	}
	var res []float64
	q := []*TreeNode{root}
	for len(q) > 0 {
		sumLevel := 0
		sizeLevel := len(q)
		for i := 0; i < sizeLevel; i++ {
			// 当前层内所有节点求和
			cur := q[0]
			q = q[1:]
			sumLevel += cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		// 计算并保存当前层的平均值
		res = append(res, float64(sumLevel)/float64(sizeLevel))
	}

	return res
}
