package 双指针解决数组问题

import (
	"reflect"
	"testing"
)

func Test_Rotate(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		K            int
		ExpectResult []int
	}{
		0: {
			Name:         "case 1",
			Nums:         []int{1, 2, 3, 4, 5, 6, 7},
			K:            3,
			ExpectResult: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for i, testCase := range testCaseList {
		nums, k := testCase.Nums, testCase.K
		expectResult := testCase.ExpectResult
		rotate(nums, k)
		if !reflect.DeepEqual(expectResult, nums) {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult =%v", i, expectResult, nums)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// rotate 使用一个新数组记录新的元素位置
func rotate(nums []int, k int) {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newIndex := (i + k) % len(nums)
		res[newIndex] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}
}

// rotateV2 多次旋转法
func rotateV2(nums []int, k int) {
	// 旋转k%n次
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// reverse 反转数组 [1,2,3,4,5]
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
