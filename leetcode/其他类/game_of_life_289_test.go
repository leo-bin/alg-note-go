package 其他类

import (
	"reflect"
	"testing"
)

// 生命游戏，289，mid
// https://leetcode.cn/problems/game-of-life/description/
// 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞
// 每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）
// 每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
// 1、如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 2、如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 3、如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 4、如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
// 下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的
// 其中细胞的出生和死亡是 同时 发生的
// 给你 m x n 网格面板 board 的当前状态，返回下一个状态
// 给定当前 board 的状态，更新 board 到下一个状态
//
// 示例1:
// 输入：board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// 输出：[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
//
// 示例2:
// 输入：board = [[1,1],[1,0]]
// 输出：[[1,1],[1,1]]
func Test_GameOfLife(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Board        [][]int
		ExpectResult [][]int
	}{
		0: {
			Name: "case 1",
			Board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			ExpectResult: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
	}

	for i, testCase := range testCaseList {
		board := testCase.Board
		expect := testCase.ExpectResult

		gameOfLife(board)
		actual := board

		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// gameOfLife dfs
func gameOfLife(board [][]int) {
	// 1 = 原来活，现在也活
	// 0 = 原来死，现在也死
	// -1 = 原来活，现在死
	// 2 = 原来死，现在活

	// base case
	if len(board) <= 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 上下左右以及对角线的8个位置同时判断最终状态
			tmp := isAlive(i-1, j, m, n, board) + isAlive(i, j-1, m, n, board) +
				isAlive(i+1, j, m, n, board) + isAlive(i, j+1, m, n, board) +
				isAlive(i-1, j-1, m, n, board) + isAlive(i+1, j+1, m, n, board) +
				isAlive(i-1, j+1, m, n, board) + isAlive(i+1, j-1, m, n, board)
			if tmp < 2 && board[i][j] == 1 {
				board[i][j] = -1
			} else if tmp > 3 && board[i][j] == 1 {
				board[i][j] = -1
			} else if tmp == 3 && board[i][j] == 0 {
				board[i][j] = 2
			} else if (tmp == 2 || tmp == 3) && board[i][j] == 1 {
				board[i][j] = 1
			}
		}
	}

	// 将2和-1的情况分别标记为1和0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 || board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == -1 {
				board[i][j] = 0
			}
		}
	}
}

// isAlive 判断当前位置的细胞状态，并标记状态所代表的枚举值
func isAlive(i, j, m, n int, board [][]int) int {
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	if board[i][j] == 1 || board[i][j] == -1 {
		return 1
	}
	return 0
}
