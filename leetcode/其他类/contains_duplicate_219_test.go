package 其他类

import (
	"reflect"
	"testing"
)

// 存在重复元素，219，easy
// https://leetcode.cn/problems/contains-duplicate-ii/description/
// 给你一个整数数组 nums 和一个整数 k
// 判断数组中是否存在两个 不同的索引 i 和 j
// 满足 nums[i] == nums[j] 且 abs(i - j) <= k
// 如果存在，返回 true ；否则，返回 false
//
// 示例 1：
// 输入：nums = [1,2,3,1], k = 3
// 输出：true
//
// 示例 2：
// 输入：nums = [1,0,1,1], k = 1
// 输出：true
//
// 示例 3：
// 输入：nums = [1,2,3,1,2,3], k = 2
// 输出：false
//
// 提示：
// 1 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
// 0 <= k <= 10^5

func Test_ContainsNearbyDuplicate(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		K            int
		ExpectResult bool
	}{
		0: {
			Name:         "case 1",
			Nums:         []int{1, 2, 3, 1},
			K:            3,
			ExpectResult: true,
		},
		1: {
			Name:         "case 2",
			Nums:         []int{1, 0, 1, 1},
			K:            1,
			ExpectResult: true,
		},
		2: {
			Name:         "case 3",
			Nums:         []int{1, 2, 3, 1, 2, 3},
			K:            2,
			ExpectResult: false,
		},
		3: {
			Name:         "case 4",
			Nums:         []int{99, 99},
			K:            2,
			ExpectResult: true,
		},
	}

	for i, testCase := range testCaseList {
		nums, k := testCase.Nums, testCase.K
		expect := testCase.ExpectResult
		actual := containsNearbyDuplicate(nums, k)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// containsNearbyDuplicate hash
func containsNearbyDuplicate(nums []int, k int) bool {
	// 1、build hash table
	hashTable := make(map[int]int, len(nums))
	// 2、find table and judge index if correct
	for i := 0; i < len(nums); i++ {
		curNum := nums[i]
		j, ok := hashTable[curNum]
		// 不存在记录下数字和所在位置
		if !ok {
			hashTable[curNum] = i
			continue
		}
		r := abs(j - i)
		// 符合条件直接返回
		if r <= k {
			return true
		} else {
			// 离k较远，找一个近的进行位置替换
			hashTable[curNum] = i
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
