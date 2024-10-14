package 排序

import (
	"reflect"
	"testing"
)

// 排序数组，912，mid
// https://leetcode.cn/problems/sort-an-array/description/
// 给你一个整数数组nums，请你将该数组升序排列
// 你必须在不使用任何内置函数的情况下解决问题
// 时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小
//
// 示例 1：
// 输入：nums = [5,2,3,1]
// 输出：[1,2,3,5]
//
// 示例 2：
// 输入：nums = [5,1,1,2,0,0]
// 输出：[0,0,1,1,2,5]

func Test_SortArray(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		ExpectResult []int
	}{
		0: {
			Name:         "case 1",
			Nums:         []int{5, 2, 3, 1},
			ExpectResult: []int{1, 2, 3, 5},
		},
		1: {
			Name:         "case 2",
			Nums:         []int{5, 1, 1, 2, 0, 0},
			ExpectResult: []int{0, 0, 1, 1, 2, 5},
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		expectResult := testCase.ExpectResult
		actualResult := sortArray(nums)
		if !reflect.DeepEqual(expectResult, actualResult) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

func sortArray(nums []int) []int {
	// base case
	if len(nums) <= 1 {
		return nums
	}
	// 随机打乱原数组的顺序，防止出现极端耗时case，比如[6,5,4,3,2,1] ,这种顺序在快排下的时间复杂度会退化成O（n*n）
	//shuffleQuickSort(nums)
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// quickSort 快排，时间复杂度O（nlogn），空间复杂度是O（logn）平衡二叉树的树高
// 思路：
// 1.分治思想，先把一个元素在数组[i,j]中的相对顺序找到，假设是p
// 2.找到p后，在对i-p之间和p到j之间的元素分别执行上述排序
// 3.最后整体就变的有序了
func quickSort(nums []int, l, r int) {
	// 递归结束
	if l >= r {
		return
	}
	p := partitionQuickSort(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partitionQuickSort(nums []int, l, r int) int {
	pos := l + (r-l)/2
	nums[l], nums[pos] = nums[pos], nums[l]
	p := nums[l]
	i, j := l+1, r
	for {
		// 从左往右找
		for i <= j && nums[i] < p {
			i++
		}
		// 从右往左找
		for i <= j && nums[j] > p {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	// p的位置符合前面的小于p，后面的要大于p
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
