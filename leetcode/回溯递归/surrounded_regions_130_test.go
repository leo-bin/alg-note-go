package 回溯递归

import (
	"reflect"
	"testing"
)

// 被围绕的区域，130，mid
// https://leetcode.cn/problems/surrounded-regions/
// 给你一个mxn的矩阵board ，由若干字符 'X' 和 'O' 组成，捕获所有被围绕的区域：
// 连接：一个单元格与水平或垂直方向上相邻的单元格连接
// 区域：连接所有 'O' 的单元格来形成一个区域
// 围绕：如果您可以用 'X' 单元格 连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕
// 通过将输入矩阵 board 中的所有 'O' 替换为 'X' 来 捕获被围绕的区域
//
// 示例 1：
// 输入：board = [["X","X","X","X"],
//               ["X","O","O","X"],
//               ["X","X","O","X"],
//               ["X","O","X","X"]]
// 输出：[["X","X","X","X"],
//       ["X","X","X","X"],
//       ["X","X","X","X"],
//       ["X","O","X","X"]]
// 解释：在上图中，底部区域没有被捕获，因为它在board边缘并且不能被围绕
//
// 示例 2：
// 输入：board = [["X"]]
// 输出：[["X"]]
//
// board =
//  [["O","O","O"],
//   ["O","O","O"],
//   ["O","O","O"]]
//
//  [["O","O","O"],
//   ["O","X","O"],
//   ["O","O","O"]]
// expect
// [["O","O","O"],
//  ["O","O","O"],
//  ["O","O","O"]]

func Test_SurroundedRegions(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Board        [][]byte
		ExpectResult [][]byte
	}{
		0: {
			Name: "case 1",
			Board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			ExpectResult: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		1: {
			Name: "case 2",
			Board: [][]byte{
				{'X'},
			},
			ExpectResult: [][]byte{
				{'X'},
			},
		},
	}

	for i, testCase := range testCaseList {
		board := testCase.Board
		expect := testCase.ExpectResult
		solve(board)
		actual := board
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed,expect=%v,actual=%v", i, expect, actual)
		}
	}
}

// solve dfs
// 思路：
// 1、把区域里的'O'识别出来，并全部标记为x
// 2、边缘的'O'不算区域，需要特殊处理
func solve(board [][]byte) {
	// base case
	if len(board) <= 0 {
		return
	}
	m, n := len(board), len(board[0])
	// 固定列，从上往下
	for i := 0; i < m; i++ {
		// 左边
		dfs(i, 0, m, n, board)
		// 右边
		dfs(i, n-1, m, n, board)
	}
	// 固定行，从左往右
	for i := 0; i < n; i++ {
		// 上边
		dfs(0, i, m, n, board)
		// 下边
		dfs(m-1, i, m, n, board)
	}
	// 将'A'还原为'O'，'O'换成'X'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// dfs 从指定的i和j开始在board里查找'O'并标记为'A'，并以i和j开始的：左、右、上、下四个方向开始继续查找
func dfs(i, j, m, n int, board [][]byte) {
	// 递归结束条件
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	dfs(i-1, j, m, n, board)
	dfs(i+1, j, m, n, board)
	dfs(i, j-1, m, n, board)
	dfs(i, j+1, m, n, board)
}
