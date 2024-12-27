package 回溯递归

import (
	"reflect"
	"testing"
)

// 课程表II，210，mid
// https://leetcode.cn/problems/course-schedule-ii/description/
// 现在你总共有 numCourses 门课需要选，记为0到numCourses - 1
// 给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi
// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1]
// 返回你为了学完所有课程所安排的学习顺序
// 可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组
//
// 示例 1：
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：[0,1]
// 解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1]
//
// 示例 2：
// 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// 输出：[0,2,1,3]
// 解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后
// 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3]
//
// 示例 3：
// 输入：numCourses = 1, prerequisites = []
// 输出：[0]

func Test_CoursesII(t *testing.T) {
	testCaseList := []struct {
		Name          string
		NumCourses    int
		Prerequisites [][]int
		ExpectResult  []int
	}{
		0: {
			Name:       "case 1",
			NumCourses: 2,
			Prerequisites: [][]int{
				{1, 0},
			},
			ExpectResult: []int{0, 1},
		},
		1: {
			Name:       "case 2",
			NumCourses: 4,
			Prerequisites: [][]int{
				{1, 0},
				{2, 0},
				{3, 1},
				{3, 2},
			},
			ExpectResult: []int{0, 2, 1, 3},
		},
		2: {
			Name:          "case 3",
			NumCourses:    1,
			Prerequisites: [][]int{},
			ExpectResult:  []int{0},
		},
		3: {
			Name:       "case 4",
			NumCourses: 2,
			Prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			ExpectResult: []int{},
		},
	}
	for i, testCase := range testCaseList {
		numCourses, prerequisites := testCase.NumCourses, testCase.Prerequisites
		expect := testCase.ExpectResult
		actual := findOrder(numCourses, prerequisites)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// findOrder dfs + 拓扑排序
// 思路：
// 1.按照图后续遍历中后反转就是拓扑排序的结果
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		curPath   = make([]bool, numCourses)
		visited   = make([]bool, numCourses)
		canFinish = true
		graph     = buildGraphII(numCourses, prerequisites)
		order     = make([]int, 0)
		dfs       = func(i int) {}
	)
	dfs = func(i int) {
		// 有循环说明无法完成
		if curPath[i] {
			canFinish = false
			return
		}
		// 提前结束
		if visited[i] || !canFinish {
			return
		}
		visited[i] = true
		curPath[i] = true
		// 继续找当前节点依赖的节点
		for _, g := range graph[i] {
			dfs(g)
		}
		curPath[i] = false
		order = append(order, i)
	}
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	if !canFinish {
		return []int{}
	}
	// 反转order就是真正上课的顺序
	reverse(order)
	return order
}

func buildGraphII(size int, prerequisites [][]int) [][]int {
	graph := make([][]int, size)
	for _, pre := range prerequisites {
		from, to := pre[1], pre[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
