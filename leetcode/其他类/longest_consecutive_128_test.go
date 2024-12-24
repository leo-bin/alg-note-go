package 其他类

import (
	"reflect"
	"testing"
)

// 最长连续序列，128，mid
// 给定一个未排序的整数数组 nums
// 找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题
// https://leetcode.cn/problems/longest-consecutive-sequence/
// 示例 1：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4],它的长度为 4
//
// 示例 2：
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9

func Test_LongestConsecutive(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Nums:         []int{100, 4, 200, 1, 3, 2},
			ExpectResult: 4,
		},
		1: {
			Name:         "case 2",
			Nums:         []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			ExpectResult: 9,
		},
		2: {
			Name:         "case 3",
			Nums:         []int{7, 1, 8, 4, 3, 2},
			ExpectResult: 4,
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		expect := testCase.ExpectResult
		actual := longestConsecutive(nums)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual =%v", i, expect, actual)
		} else {
			t.Logf("test caee %v,passed", i)
		}
	}
}

// longestConsecutive hash
func longestConsecutive(nums []int) int {
	n := len(nums)
	// base case
	if n <= 0 {
		return 0
	}
	maxLen := 1
	record := make(map[int]struct{}, n)
	// 初始化长度hash表
	for i := 0; i < n; i++ {
		record[nums[i]] = struct{}{}
	}
	for num := range record {
		// 找序列的开头元素，不是开头元素的直接跳过
		if _, ok := record[num-1]; ok {
			continue
		}
		// 从当前元素开始不断+1，找到所有的连续序列
		tmp := 1
		next := num + 1
		for {
			if _, ok2 := record[next]; !ok2 {
				break
			}
			tmp++
			next++
		}
		if maxLen < tmp {
			maxLen = tmp
		}
	}
	return maxLen
}
