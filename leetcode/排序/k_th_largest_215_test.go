package 排序

import (
	"testing"
)

// 数组中最大的第k个元素，215，mid
// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题
//
// 示例 1:
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
//
// 示例 2:
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4

func Test_KthLargest(t *testing.T) {

}

// kthLargest 快速选择算法，时间复杂度O（n）
// 思想：
// 1.首先找第k大的元素就等同于找第n-k个元素
// 2.快速排序在找p值时（partition）的过程中本身就是在找p在数组中的排名rank
// 3.根据这个我们只需要知道每次rank的结果和n-k进行比较进行
// 4.如果rank<p，说明p不够用了，得扩大查找范围，继续往右找
// 5.如果rank>p，说明得往左继续找
func kthLargest(nums []int, k int) int {
	n := len(nums)
	// 第k大等同于第n-k个元素
	k = n - k
	l, r := 0, n-1
	for l <= r {
		p := partitionKth(nums, l, r)
		if p < k {
			l = p + 1
		} else if p > k {
			r = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func partitionKth(nums []int, l, r int) int {
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
