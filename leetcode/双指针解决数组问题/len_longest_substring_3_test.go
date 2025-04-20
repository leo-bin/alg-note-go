package 双指针解决数组问题

import "testing"

// 无重复字符的最长子串，3，mid
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
//
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1
//
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
// 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串
func TestLengthOfLongestSubstring(t *testing.T) {

}

// hash+滑动窗口
// 主要，子串要求是连续的，要删除只能从头删除元素
func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	i, j := 0, 0
	maxLen := 0
	for j < len(s) {
		// 记录当前的元素
		hash[s[j]]++
		for hash[s[j]] > 1 {
			// 如果出现重复元素需要收缩最左侧窗口
			hash[s[i]]--
			i++
		}
		// 更新长度
		maxLen = maxInt(maxLen, j-i+1)
		j++
	}
	return maxLen
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
