package 双指针解决数组问题

import (
	"sort"
	"testing"
)

// 安排会议室i，252，easy
// https://leetcode.cn/problems/meeting-rooms/
// 给定一个会议时间安排的数组 intervals
// 每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi]
// 请你判断一个人是否能够参加这里面的全部会议
//
// 示例 1：
// 输入：intervals = [[0,30],[5,10],[15,20]]
// 输出：false
//
// 示例 2：
// 输入：intervals = [[7,10],[2,4]]
// 输出：true

func Test_MeetingI(t *testing.T) {

}

// canAttendMeeting 双指针+排序
// 思路：
// 1.收集会议室开始和结束时间记为start和end数组，并按自然时间序列排序（从小到大）
// 2.使用两个指针i、j分别从start和end数组开始遍历，模拟一条自然时间线的走向
// 3.使用count记录同一时刻一共需要的会议室数量
// 4.如果start[i]<end[j]说明两个会议室冲突，需要加一个会议室，count+1同时i往后移，否则count-1同时j+往后移
// 5.如果count>=2，说明有会议冲突，1一个人无法同时参加
func canAttendMeeting(intervals [][]int) bool {
	// base case
	n := len(intervals)
	if n <= 0 {
		return true
	}
	// 1.收集会议室开始和结束时间记为start和end数组
	start, end := make([]int, n), make([]int, n)
	for i, interval := range intervals {
		start[i] = interval[0]
		end[i] = interval[1]
	}
	// 2.按自然时间序列排序（从小到大）
	sort.Ints(start)
	sort.Ints(end)
	// 3.模拟一条自然时间线的走向
	i, j, count, maxCnt := 0, 0, 0, 0
	for i < n && j < n {
		if start[i] < end[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		maxCnt = maxMeeting(maxCnt, count)
	}
	// 4.判断最多是否需要1个会议室
	if maxCnt == 1 {
		return true
	} else {
		return false
	}
}

func maxMeeting(x, y int) int {
	if x > y {
		return x
	}
	return y
}
