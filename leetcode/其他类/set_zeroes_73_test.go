package 其他类

import (
	"math"
	"reflect"
	"testing"
)

// 矩阵置零，73，mid
// https://leetcode.cn/problems/set-matrix-zeroes/description/
// 给定一个m x n的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为0
// 请使用 原地 算法
//
// 示例1:
// 输入：[[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
// 示例2:
// 输入：[[-1],[2],[3]]
// 输出：[[-1],[2],[3]]

func Test_SetZeroes(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Matrix       [][]int
		ExpectResult [][]int
	}{
		0: {
			Name: "case 1",
			Matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			ExpectResult: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		1: {
			Name: "case 2",
			Matrix: [][]int{
				{-1},
				{2},
				{3},
			},
			ExpectResult: [][]int{
				{-1},
				{2},
				{3},
			},
		},
	}

	for i, testCase := range testCaseList {
		matrix := testCase.Matrix
		expect := testCase.ExpectResult
		setZeroes(matrix)
		actual := matrix
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// setZeroes 原地修改
func setZeroes(matrix [][]int) {
	// base case
	if len(matrix) <= 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				// i、j所在行标记为-1
				for z := 0; z < n; z++ {
					if matrix[i][z] != 0 {
						matrix[i][z] = math.MinInt
					}
				}
				// i、j所在列标记为-1
				for z := 0; z < m; z++ {
					if matrix[z][j] != 0 {
						matrix[z][j] = math.MinInt
					}
				}
			}
		}
	}
	// -1置为0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == math.MinInt {
				matrix[i][j] = 0
			}
		}
	}
	return
}
