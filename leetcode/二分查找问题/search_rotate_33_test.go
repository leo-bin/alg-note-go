package 二分查找问题

// 搜索旋转数组，mid
// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4

// 思路： 一次二分查找
// 1.原数组由两个升序数组组成，我们需要先找到target在哪个数组
// 2.借鉴二分查找的核心思想，先定位mid，判断mid在哪个数组中
// 3.进入那个数组中机洗二分查找具体元素
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		}
		// 判断mid在哪个数组中
		if nums[mid] < nums[i] {
			// mid比前面的数还要小，说明mid在后边的有序数组中
			if nums[mid] < target && nums[j] >= target {
				// 往更大的范围收缩
				i = mid + 1
			} else {
				j = mid - 1
			}
		} else {
			if nums[i] <= target && nums[mid] > target {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}
	return -1
}

// 思路：
// 1.找逆序的地方，把数组分为两段有序的数组，两次二分查找即可

// [4,5,6,7,0,1,2]
// [3,1],1
func searchV2(nums []int, target int) int {
	// 1.找逆序的下标
	i := 0
	for ; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			break
		}
	}
	// 2.分别在两个数组中进行二分查找
	find1 := binSearch(nums, 0, i, target)
	find2 := binSearch(nums, i+1, len(nums)-1, target)
	if find1 != -1 {
		return find1
	}
	if find2 != -1 {
		return find2
	}
	return -1
}

// 二分查找
func binSearch(nums []int, i, j, target int) int {
	// base case
	if i == j {
		if nums[i] == target {
			return i
		} else {
			return -1
		}
	}
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
