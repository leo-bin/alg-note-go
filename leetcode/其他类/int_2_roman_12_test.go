package 其他类

import "testing"

// 整形转罗马数字，12，mid
// https://leetcode.cn/problems/integer-to-roman/description/
// 七个不同的符号代表罗马数字，其值如下：
// 符号	值
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
// 罗马数字是通过添加从最高到最低的小数位值的转换而形成的。将小数位值转换为罗马数字有以下规则：
// 如果该值不是以 4 或 9 开头，请选择可以从输入中减去的最大值的符号，将该符号附加到结果，减去其值，然后将其余部分转换为罗马数字。
// 如果该值以 4 或 9 开头，使用 减法形式，表示从以下符号中减去一个符号
// 例如 4 是 5 (V) 减 1 (I): IV ，9 是 10 (X) 减 1 (I)：IX
// 仅使用以下减法形式：4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)
// 只有 10 的次方（I, X, C, M）最多可以连续附加 3 次以代表 10 的倍数
// 你不能多次附加 5 (V)，50 (L) 或 500 (D)。如果需要将符号附加4次，请使用 减法形式
// 给定一个整数，将其转换为罗马数字
//
// 示例 1：
// 输入：num = 3749
// 输出： "MMMDCCXLIX"
// 解释：
// 3000 = MMM 由于 1000 (M) + 1000 (M) + 1000 (M)
// 700 = DCC 由于 500 (D) + 100 (C) + 100 (C)
//  40 = XL 由于 50 (L) 减 10 (X)
//   9 = IX 由于 10 (X) 减 1 (I)
// 注意：49 不是 50 (L) 减 1 (I) 因为转换是基于小数位
//
// 示例 2：
// 输入：num = 58
// 输出："LVIII"
// 解释：
// 50 = L
// 8 = VIII
//
// 示例 3：
// 输入：num = 1994
// 输出："MCMXCIV"
// 解释：
// 1000 = M
// 900 = CM
//  90 = XC
//   4 = IV
//
// 提示：
// 1 <= num <= 3999

func Test_Int2Roman(t *testing.T) {

}

var romanIntValueList = []struct {
	Roman string
	Value int
}{
	0:  {"M", 1000},
	1:  {"CM", 900},
	2:  {"D", 500},
	3:  {"CD", 400},
	4:  {"C", 100},
	5:  {"XC", 90},
	6:  {"L", 50},
	7:  {"XL", 40},
	8:  {"X", 10},
	9:  {"IX", 9},
	10: {"V", 5},
	11: {"IV", 4},
	12: {"I", 1},
}

// intToRoman 模拟
// 思路：
// 1.根据规则，大的罗马字符在前面，我们只需要根据所有已知罗马字符的值减去当前的num，最后拼接起来的就是完整的字符
func intToRoman(num int) string {
	var res []byte
	for _, romanValue := range romanIntValueList {
		for num >= romanValue.Value {
			num -= romanValue.Value
			res = append(res, romanValue.Roman...)
		}
		if num <= 0 {
			break
		}
	}
	return string(res)
}
