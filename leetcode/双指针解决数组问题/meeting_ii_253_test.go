package 双指针解决数组问题

import (
	"sort"
	"testing"
)

// 安排会议室ii，252，mid
// https://leetcode.cn/problems/meeting-rooms-ii/description/
// 给你一个会议时间安排的数组 intervals
// 每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi]
// 返回 所需会议室的最小数量
//
// 示例 1：
// 输入：intervals = [[0,30],[5,10],[15,20]]
// 输出：2
//
// 示例 2：
// 输入：intervals = [[7,10],[2,4]]
// 输出：1

func Test_MeetingII(t *testing.T) {

}

// minMeetingRoom 排序+双指针模拟
// 思路：
// 1.收集会议室开始和结束时间记为start和end数组，并按自然时间序列排序（从小到大）
// 2.使用两个指针i、j分别从start和end数组开始遍历，模拟一条自然时间线的走向
// 3.使用count记录同一时刻一共需要的会议室数量,maxRes记录最大的count值
// 4.如果start[i]<end[j]说明两个会议室冲突，需要加一个会议室，count+1同时i往后移，否则count-1同时j+往后移
func minMeetingRoom(intervals [][]int) int {
	// base case
	n := len(intervals)
	if n <= 0 {
		return 0
	}
	// 1.收集会议室开始和结束时间记为start和end数组
	start, end := make([]int, n), make([]int, n)
	for i, interval := range intervals {
		start[i] = interval[0]
		end[i] = interval[1]
	}
	// 2.排序
	sort.Ints(start)
	sort.Ints(end)
	// 3.模拟一条自然时间线的走向,记录count
	i, j, count, maxRes := 0, 0, 0, 0
	for i < n && j < n {
		if start[i] < end[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		maxRes = maxMeetingII(maxRes, count)
	}
	// 4.返回最多需要几间会议室
	return maxRes
}

func maxMeetingII(x, y int) int {
	if x > y {
		return x
	}
	return y
}
