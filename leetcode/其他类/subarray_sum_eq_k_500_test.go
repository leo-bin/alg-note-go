package 其他类

// 和为K的子数组，500，mid
// https://leetcode.cn/problems/subarray-sum-equals-k/
// 给你一个整数数组 nums 和一个整数 k
// 请你统计并返回 该数组中和为 k 的子数组的个数
// 子数组是数组中元素的连续非空序列
//
// 示例 1：
// 输入：nums = [1,1,1], k = 2
// 输出：2
//
// 示例 2：
// 输入：nums = [1,2,3], k = 3
// 输出：2

// 思路： 前缀和+hash表
// 因为我们需要找到和为k的连续子数组的个数
// 通过计算前缀和，我们可以将问题转化为求解两个前缀和之差等于k的情况
// 假设数组的前缀和数组为prefixSum，其中prefixSum[i]表示从数组起始位置到第i个位置的元素之和
// 那么对于任意的两个下标i和j（i < j），如果prefixSum[j] - prefixSum[i] = k，即从第i个位置到第j个位置的元素之和等于k
// 那么说明从第i+1个位置到第j个位置的连续子数组的和为k
// 通过遍历数组，计算每个位置的前缀和，并使用一个哈希表来存储每个前缀和出现的次数
// 在遍历的过程中，我们检查是否存在prefixSum[j] - k的前缀和
// 如果存在，说明从某个位置到当前位置的连续子数组的和为k，我们将对应的次数累加到结果中

// map[0]=1 -> []出现1次
// map[1]=1 -> [1]出现1次
// map[3]=1 -> [1,2]出现1次
// map[6]=1 -> [1,2,3]出现1次
func subarraySum(nums []int, k int) int {
	cnt := 0
	sum := 0
	hash := make(map[int]int)
	hash[0] = 1
	for i := 0; i < len(nums); i++ {
		// 1.计算前缀和
		sum += nums[i]
		// 2.通过前缀和判断某个连续子数组是否出现
		if v, ok := hash[sum-k]; ok {
			cnt += v
		}
		// 更新前缀和出现次数
		hash[sum] += 1
	}
	return cnt
}
