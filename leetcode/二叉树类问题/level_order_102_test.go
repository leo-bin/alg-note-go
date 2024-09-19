package 二叉树类问题

import (
	"reflect"
	"testing"
)

// 二叉树的层序遍历，leetcode-102，mid
// 链接：https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
// 给你二叉树的根节点 root ，返回其节点值的层序遍历
// 逐层地，从左到右访问所有节点
//
// 示例1:
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
// 示例2:
// 输入：root = [1]
// 输出：[[1]]
//
// 示例 3：
// 输入：root = []
// 输出：[]

func Test_LevelOrder(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Root         *TreeNode
		ExpectResult [][]int
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
			ExpectResult: [][]int{{3}, {9, 20}, {15, 7}},
		},
		1: {
			Name: "case 2",
			Root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			ExpectResult: [][]int{{1}},
		},
		2: {
			Name:         "case 3",
			Root:         nil,
			ExpectResult: [][]int{},
		},
	}

	for i, testCase := range testCaseList {
		root := testCase.Root
		expectResult := testCase.ExpectResult
		actualResult := levelOrder(root)
		if !reflect.DeepEqual(actualResult, expectResult) {
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

// levelOrder 层序遍历
// 思路：
// 1.标准的层序遍历框架，注意在哪个地方收集数据即可
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		var levelRes []int
		curSize := len(q)
		for i := 0; i < curSize; i++ {
			// 收集当前层的所有元素
			cur := q[0]
			q = q[1:]
			levelRes = append(levelRes, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		// 收集每一层的所有元素
		res = append(res, levelRes)
	}

	return res
}
