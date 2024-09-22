package 二叉树类问题

import "testing"

func Test_ConnectTree(t *testing.T) {
	testCaseList := []struct {
		Name string
		Root *Node
	}{
		0: {
			Name: "case 1",
			Root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val:  3,
					Left: nil,
					Right: &Node{
						Val: 7,
					},
				},
			},
		},
	}

	for _, testCase := range testCaseList {
		root := testCase.Root
		connect(root)
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect bfs，直接连接每一层的所有元素
func connect(root *Node) *Node {
	// base case
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		var levelRes []*Node
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			levelRes = append(levelRes, cur)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		// 直接连接每一层的所有元素
		for i := 1; i < len(levelRes); i++ {
			pre, cur := levelRes[i-1], levelRes[i]
			pre.Next = cur
		}
	}
	return root
}
