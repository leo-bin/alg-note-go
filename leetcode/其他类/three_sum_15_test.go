package 其他类

import (
	"reflect"
	"sort"
	"testing"
)

// 三数之和，15，mid
// https://leetcode.cn/problems/3sum/
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
// 满足 i != j、i != k 且 j != k
// 同时还满足 nums[i] + nums[j] + nums[k] == 0
// 请你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
//
// 示例 2：
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0 。
//
// 示例 3：
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0

func Test_ThreeSum(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		ExpectResult [][]int
	}{
		0: {"case 1",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				0: {-1, -1, 2},
				1: {-1, 0, 1},
			},
		},
		1: {"case 2",
			[]int{0, 1, 1},
			[][]int{
				0: {},
			},
		},
		2: {"case 3",
			[]int{0, 0, 0},
			[][]int{
				0: {0, 0, 0},
			},
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		expectResult := testCase.ExpectResult
		actualResult := threeSum(nums)
		if !reflect.DeepEqual(expectResult, actualResult) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// threeSum 排序+双指针
// 思路：
// 1.求三数之和，我们只需要固定每一个元素，从i+1到end范围内，找两数之和=-nums[i]的数即可
// 2.需要注意的是中间会遇到第一个元素为重复的case，需要原地进行去重
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	// 1.先排序
	sort.Ints(nums)
	// 2.固定一个元素，找剩余元素的两数之和
	for i := 0; i < len(nums); i++ {
		tmpResult := twoSumTarget(nums, i+1, -nums[i])
		for _, tmp := range tmpResult {
			tmp = append(tmp, nums[i])
			res = append(res, tmp)
		}
		// 防止第一个元素重复
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// twoSumTarget 在指定范围内找 两数只和=target的二元组集合
func twoSumTarget(nums []int, start int, target int) [][]int {
	l, r := start, len(nums)-1
	res := make([][]int, 0)
	for l < r {
		sum := nums[l] + nums[r]
		left, right := nums[l], nums[r]
		if sum == target {
			res = append(res, []int{left, right})
			// 防止重复组合出现
			for l < r && nums[l] == left {
				l++
			}
			for l < r && nums[r] == right {
				r--
			}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return res
}
