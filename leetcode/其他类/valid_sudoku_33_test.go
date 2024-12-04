package 其他类

import (
	"reflect"
	"testing"
)

// 有效的数独，33，mid
// https://leetcode.cn/problems/valid-sudoku/description/
// 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可
//
// 数字 1-9 在每一行只能出现一次
// 数字 1-9 在每一列只能出现一次
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次
// 注意：
// 一个有效的数独（部分已被填充）不一定是可解的
// 只需要根据以上规则，验证已经填入的数字是否有效即可
// 空白格用 '.' 表示

func Test_ValidSudoku(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Board        [][]byte
		ExpectResult bool
	}{
		0: {
			Name: "case 1",
			Board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			ExpectResult: true,
		},
	}

	for i, testCase := range testCaseList {
		board := testCase.Board
		expect := testCase.ExpectResult
		actual := isValidSuDoku(board)
		if reflect.DeepEqual(expect, actual) {
			t.Logf("test case %v,passed", i)
		} else {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		}
	}
}

// isValidSuDoKu 打表
func isValidSuDoku(board [][]byte) bool {
	var rowTable, colTable [9][9]int
	var boxsTable [3][3][9]int
	for i, row := range board {
		for j, num := range row {
			if num == '.' {
				continue
			}
			index := num - '1'
			rowTable[i][index]++
			colTable[j][index]++
			boxsTable[i/3][j/3][index]++
			if rowTable[i][index] > 1 || colTable[j][index] > 1 || boxsTable[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
