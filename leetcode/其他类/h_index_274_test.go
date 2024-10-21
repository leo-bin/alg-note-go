package 其他类

import (
	"sort"
	"testing"
)

// h指数，274，mid
// https://leetcode.cn/problems/h-index/
// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数
// 计算并返回该研究者的 h 指数
// 根据维基百科上 h 指数的定义：
// h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文
// 并且 至少 有 h 篇论文被引用次数大于等于 h
// 如果 h 有多种可能的值，h 指数 是其中最大的那个。
//
// 示例 1：
// 输入：citations = [3,0,6,1,5]
// 输出：3
// 解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次
// 由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3
//
// 示例 2：
// 输入：citations = [1,3,1]
// 输出：1

func Test_HIndex(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Citations    []int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Citations:    []int{3, 0, 6, 1, 5},
			ExpectResult: 3,
		}, 1: {
			Name:         "case 2",
			Citations:    []int{1, 3, 1},
			ExpectResult: 1,
		},
	}

	for i, testCase := range testCaseList {
		c := testCase.Citations
		expectResult := testCase.ExpectResult
		actualResult := hIndex(c)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// hIndex h指数 排序+正序递推
// 思路：
// 1.首先h指数的取值范围 = {0 <= h <= 论文个数}
// 2.以示例一的数据分析，把论文引用次数从小到大排序，得到c=[0,1,3,5,6]
// 3.逐步分析h的可能取值：
// h=0表示，至少有0篇论文，至少被引用了0次【成立】
// h=1表示，至少有1篇论文，至少被引用了1次【成立】
// h=2表示，至少有2篇论文，至少被引用了2次【成立】
// h=3表示，至少有3篇论文，至少被引用了3次【成立】答案
// h=4表示，至少有4篇论文，至少被引用了4次【不成立】
// h=5。。。。【不成立，h=4都不成立，h=5肯定不成立】
// 从最小值开始，如果最小引用次数比5大（总论文个数），那不用想了，说明全部引用次数都比5大，结果就是5
// 倒数第二小值，此时的最大答案是4（去掉最小值），如果倒数第二小值比4大，同样的，剩下的都比4大，结果就是4
// 倒数第三小值，此时的最大答案是3，同样的，只要它比3大，剩下的都比3大，结果就是3
// 以此类推
func hIndex(citations []int) int {
	sort.Ints(citations)
	maxH := len(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= maxH {
			return maxH
		} else {
			maxH--
		}
	}
	return maxH
}
