package 其他类

import (
	"reflect"
	"testing"
)

// 旋转图像，48,mid
// https://leetcode.cn/problems/rotate-image/description/
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像
// 请你将图像顺时针旋转 90 度
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵
// 请不要 使用另一个矩阵来旋转图像

func Test_Rotate(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Matrix       [][]int
		ExpectResult [][]int
	}{
		0: {
			Name: "case 1",
			Matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			ExpectResult: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}
	for i, testCase := range testCaseList {
		matrix := testCase.Matrix
		expect := testCase.ExpectResult
		// run
		rotate(matrix)
		actual := matrix
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test caese %v,passed", i)
		}
	}
}

// rotate
func rotate(matrix [][]int) {
	n := len(matrix)
	// 沿对角线对元素进行镜像
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 翻转每一行
	for _, rows := range matrix {
		reverse(rows)
	}
	return
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
