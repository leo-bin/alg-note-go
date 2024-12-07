package 其他类

import (
	"reflect"
	"testing"
)

// 螺旋矩阵，54，mid
// https://leetcode.cn/problems/spiral-matrix/description/
// 给你一个 m 行 n 列的矩阵 matrix
// 请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

func Test_SpiralOrder(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Matrix       [][]int
		ExpectResult []int
	}{
		0: {
			Name: "case 1",
			Matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			ExpectResult: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}

	for i, testCase := range testCaseList {
		matrix := testCase.Matrix
		expect := testCase.ExpectResult
		actual := spiralOrder(matrix)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// spiralOrder 模拟
func spiralOrder(matrix [][]int) []int {
	// base case
	if len(matrix) <= 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	left, right, up, bottom := 0, n-1, 0, m-1
	res := make([]int, 0)
	for left <= right && up <= bottom {
		// 1.left -> right
		for j := left; j <= right; j++ {
			res = append(res, matrix[up][j])
		}
		up++
		// 2.up -> bottom
		for j := up; j <= bottom; j++ {
			res = append(res, matrix[j][right])
		}
		right--
		// 3.right -> left
		if up <= bottom {
			for j := right; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
		}
		bottom--
		// 4.bottom -> up
		if left <= right {
			for j := bottom; j >= up; j-- {
				res = append(res, matrix[j][left])
			}
		}
		left++
	}
	return res
}
