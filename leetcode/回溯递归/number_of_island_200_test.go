package 回溯递归

import (
	"reflect"
	"testing"
)

// 岛屿数量，200，mid
// https://leetcode.cn/problems/number-of-islands/
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成
// 此外，你可以假设该网格的四条边均被水包围。
// 示例 1：
// 输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
// ]
// 输出：1
//
// 示例 2：
// 输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
// ]
// 输出：3

func Test_NumIslands(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Grid         [][]byte
		ExpectResult int
	}{
		0: {
			Name: "case 1",
			Grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			ExpectResult: 1,
		},
		1: {
			Name: "case 2",
			Grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			ExpectResult: 3,
		},
	}

	for i, testCase := range testCaseList {
		grid := testCase.Grid
		expect := testCase.ExpectResult
		actual := numIslands(grid)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual =%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// numIslands dfs标记法
// 思路：
// 1、从一个陆地开始不断遍历，找到相邻的所有陆地，并进行特殊标记
// 2、标记结束代表以当前陆地开始的岛屿被找到，岛屿数量+1
// 3、继续遍历直到结束
func numIslands(grid [][]byte) int {
	// base case
	if len(grid) <= 0 {
		return 0
	}
	// 初始化变量
	m, n := len(grid), len(grid[0])
	islandCnt := 0

	// 定义dfs递归函数
	var dfsFounder func(i, j int)
	dfsFounder = func(i, j int) {
		// 递归结束条件
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfsFounder(i-1, j)
		dfsFounder(i+1, j)
		dfsFounder(i, j-1)
		dfsFounder(i, j+1)
	}
	// 找岛屿并标记数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfsFounder(i, j)
				islandCnt++
			}
		}
	}
	return islandCnt
}
