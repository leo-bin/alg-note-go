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

// mergeSort 归并排序，时间复杂度O（n*logn），空间复杂度O（n）
// 思想：
// 1.先分治，直到只剩下一个元素，那么它本身就是有序的，只需要对分治的结果进行在排序合并即可
func mergeSort(nums []int, tmp []int, l, r int) {
	// 1.递归结束条件
	if l == r {
		return
	}
	// 2.对原有数组不断二分
	mid := l + (r-l)/2
	mergeSort(nums, tmp, l, mid)
	mergeSort(nums, tmp, mid+1, r)
	// 3.合并l-mid和mid-r之间的元素
	merge(nums, tmp, l, mid, r)
}

func merge(nums []int, tmp []int, l, mid, r int) {
	// 1.复制待排序范围内的元素到tmp数组（l-r）
	for i := 0; i < r; i++ {
		tmp[i] = nums[i]
	}
	// 2.双指针技巧对l-mid和mid到r进行排序
	i, j := l, mid+1
	for p := l; p <= r; p++ {
		if i == mid+1 {
			// 说明i-mid范围内的元素都已排好，剩下的mid-j直接往后放
			nums[p] = tmp[j]
			j++
		} else if j == r+1 {
			// 说明mid-j范围内的元素已排好，剩下的mid-j直接往后放
			nums[p] = tmp[i]
			i++
		} else if tmp[i] > tmp[j] {
			// 小的先放
			nums[p] = tmp[j]
			j++
		} else {
			nums[p] = tmp[i]
			i++
		}
	}
}
