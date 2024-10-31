package 其他类

import "testing"

// 最长公共前缀，14，easy
// https://leetcode.cn/problems/longest-common-prefix/description/
// 编写一个函数来查找字符串数组中的最长公共前缀
// 如果不存在公共前缀，返回空字符串 ""
//
// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
//
// 示例 2：
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀

func Test_LongestCommonPrefix(t *testing.T) {

}

// lcp 纵向比较法 时间复杂度O（m*n）
// 思路：
// 1.把所有字符看成一个二位数组
// 2.以第一个字符的列为基准，和剩余的字符串的所有列进行比较
// 3.一旦出现某两行的列不相同，说明前缀从这里就断开了，直接返回当前字符串的前面的列即可
func longestCommonPrefix(strs []string) string {
	// 行
	m := len(strs)
	if m <= 0 {
		return ""
	}
	// 列
	n := len(strs[0])
	// 固定列，判断每一行的当前列是否相同
	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			curStr, lastStr := strs[row], strs[row-1]
			if col >= len(curStr) || col >= len(lastStr) || curStr[col] != lastStr[col] {
				// 说明此时只有[0-col-1]的字符是相同的
				return strs[row][:col]
			}
		}
	}
	// 第一个字符就是所有字符的公共前缀
	return strs[0]
}
