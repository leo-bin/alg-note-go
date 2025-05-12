package 其他类

import (
	"reflect"
	"testing"
)

// 杨辉三角，118，easy
// https://leetcode.cn/problems/pascals-triangle/
// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和
// 暴力，时间复杂度O(n*n)
// 思路：把杨辉三角左对齐即可，然后用上一行的结果算就行
//
// [1]
// [1,1]
// [1,2,1]
// [1,3,3,1]
// [1,4,6,4,1]

func TestName(t *testing.T) {
	testCaseList := []struct {
		Name    string
		NumRows int
		Expect  [][]int
	}{
		0: {
			Name:    "case 1",
			NumRows: 5,
			Expect: [][]int{
				0: {1},
				1: {1, 1},
				2: {1, 2, 1},
				3: {1, 3, 3, 1},
				4: {1, 4, 6, 4, 1},
			},
		},
	}

	for i, testCase := range testCaseList {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := generate(testCase.NumRows)
			if !reflect.DeepEqual(testCase.Expect, actual) {
				t.Errorf("test case %v,not passed", i)
			} else {
				t.Logf("test case %v,passed", i)
			}
		})
	}
}

func generate(numRows int) [][]int {
	// base case
	if numRows <= 0 {
		return [][]int{}
	}

	res := make([][]int, numRows)
	// 计算每一层的元素
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		// 每一层的第一个和最后一个位置是1
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}

	return res
}
