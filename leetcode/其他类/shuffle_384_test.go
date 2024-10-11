package 其他类

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 打乱数组，384，mid
// https://leetcode.cn/problems/shuffle-an-array/description/
// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组
// 打乱后，数组的所有排列应该是 等可能 的。
// 实现 Solution class:
// Solution(int[] nums) 使用整数数组 nums 初始化对象
// int[] reset() 重设数组到它的初始状态并返回
// int[] shuffle() 返回数组随机打乱后的结果
//
// 示例 1：
// 输入
// ["Solution", "shuffle", "reset", "shuffle"]
// [[[1, 2, 3]], [], [], []]
// 输出
// [null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
// 解释
// Solution solution = new Solution([1, 2, 3]);
// solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
// solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
// solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]

func Test_Shuffle(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	s := Constructor(nums)
	res1 := s.Shuffle()
	res2 := s.Reset()
	fmt.Printf("res1=%v", res1)
	fmt.Printf("res2=%v", res2)
}

type Solution struct {
	nums []int
	rand *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums, rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (s *Solution) Reset() []int {
	return s.nums
}

// shuffle 洗牌算法
// 思路：
// 1.随机洗牌，只需要证明每一次洗牌产生的结果的概率是一致的即可
// 2.n张牌，按照全排列一共有n!种排序方法
// 3.我们只需要从每一位置开始和后面剩余的位置里的元素进行交换即可
// 4.假设nums=[1,2,3]
// nums[0]可以和后面的nums[0]、nums[1]、 nums[2]交换
// nums[1]可以和后面的nums[1]、nums[2]交换
func (s *Solution) Shuffle() []int {
	n := len(s.nums)
	cp := make([]int, n)
	copy(cp, s.nums)
	for i := 0; i < n; i++ {
		r := i + s.rand.Intn(n-i)
		// 交换位置
		cp[i], cp[r] = cp[r], cp[i]
	}
	return cp
}
