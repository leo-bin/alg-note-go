package 其他类

import "testing"

// 字符串解码，394，mid
// https://leetcode.cn/problems/decode-string/
// 给定一个经过编码的字符串，返回它解码后的字符串
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次
// 注意 k 保证为正整数
// 你可以认为输入字符串总是有效的
// 输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k
// 例如不会出现像 3a 或 2[4] 的输入
// 示例1：
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
//
// 示例2：
// 输入：s = "3[a2[c]]"
// 输出："accaccacc"

func TestDecodeString(t *testing.T) {
	s := "3[a2[c]]"
	decodeString(s)
}

// 思路：
// 1.双栈模拟，一个数字栈，一个字符栈
// 2.遇到[入栈，遇到]出栈，考虑嵌套的情况

func decodeString(s string) string {
	n := len(s)
	// 1.定义数字栈和字符栈，用来记录匹配的中间结果
	num, numStack := 0, make([]int, 0)
	str, strStack := "", make([]string, 0)
	// 2.从左往右依次入栈，遇到'['入栈，遇到']'出栈并还原字符串
	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			// 数字，存起来
			num = num*10 + int(s[i]-'0')
		} else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			// 字符，存起来
			str = str + string(s[i])
		} else if s[i] == '[' {
			// 入栈并清空数据
			numStack = append(numStack, num)
			strStack = append(strStack, str)
			num, str = 0, ""
		} else if s[i] == ']' {
			// 出栈
			cnt, item := numStack[len(numStack)-1], strStack[len(strStack)-1]
			numStack, strStack = numStack[:len(numStack)-1], strStack[:len(strStack)-1]
			// 构造真实str
			for j := 0; j < cnt; j++ {
				item += str
			}
			str = item
		}
	}
	return str
}
