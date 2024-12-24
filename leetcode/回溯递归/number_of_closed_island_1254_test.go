package 回溯递归

import (
	"reflect"
	"testing"
)

// 封闭岛屿的数量，1254，mid
// 二维矩阵 grid 由 0 （土地）和 1 （水）组成
// 岛是由最大的4个方向连通的0组成的群，封闭岛是一个完全由1包围（左、上、右、下）的岛
// 请返回封闭岛屿的数目
// https://leetcode.cn/problems/number-of-closed-islands/description/
//
// 示例 1：
// 输入：grid = [[1,1,1,1,1,1,1,0],
//              [1,0,0,0,0,1,1,0],
//              [1,0,1,0,1,1,1,0],
//              [1,0,0,0,0,1,0,1],
//              [1,1,1,1,1,1,1,0]]
// 输出：2
// 解释：
// 灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）
//
// 示例 2：
// 输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
// 输出：1
//
// 示例 3：
// 输入：grid = [[1,1,1,1,1,1,1],
//             [1,0,0,0,0,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,1,0,1,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,0,0,0,0,1],
//             [1,1,1,1,1,1,1]]
// 输出：2

func Test_CloseIsland(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Grid         [][]int
		ExpectResult int
	}{
		0: {
			Name: "case 1",
			Grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			ExpectResult: 2,
		},
		1: {
			Name: "case 2",
			Grid: [][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			},
			ExpectResult: 1,
		},
		2: {
			Name: "case 3",
			Grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			ExpectResult: 2,
		},
	}

	for i, testCase := range testCaseList {
		grid := testCase.Grid
		expect := testCase.ExpectResult
		actual := closedIsland(grid)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// closedIsland dfs标记
func closedIsland(grid [][]int) int {
	// base case
	if len(grid) <= 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	islandCnt := 0
	var dfsFounder func(i, j int) bool
	dfsFounder = func(i, j int) bool {
		// 递归结束条件
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if grid[i][j] == 1 || grid[i][j] == 2 {
			return true
		}
		grid[i][j] = 2
		r1, r2, r3, r4 := dfsFounder(i-1, j), dfsFounder(i+1, j), dfsFounder(i, j-1), dfsFounder(i, j+1)
		return r1 && r2 && r3 && r4
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && dfsFounder(i, j) {
				islandCnt++
			}
		}
	}
	return islandCnt
}
