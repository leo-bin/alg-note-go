package 回溯递归

import (
	"reflect"
	"testing"
)

// 课程表，207，mid
// https://leetcode.cn/problems/course-schedule/description/
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1
// 在选修某些课程之前需要一些先修课程
// 先修课程按数组 prerequisites 给出
// 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false
//
// 示例 1：
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true
// 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

func Test_Courses(t *testing.T) {
	testCaseList := []struct {
		Name          string
		NumCourses    int
		Prerequisites [][]int
		ExpectResult  bool
	}{
		0: {
			Name:       "case 1",
			NumCourses: 2,
			Prerequisites: [][]int{
				{1, 0},
			},
			ExpectResult: true,
		},
		1: {
			Name:       "case 2",
			NumCourses: 2,
			Prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			ExpectResult: false,
		},
	}
	for i, testCase := range testCaseList {
		numCourses, prerequisites := testCase.NumCourses, testCase.Prerequisites
		expect := testCase.ExpectResult
		actual := canFinish(numCourses, prerequisites)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// canFinish dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		curPath = make([]bool, numCourses)
		visited = make([]bool, numCourses)
		// 构造关系图
		graph  = buildGraph(numCourses, prerequisites)
		result = true
		dfs    = func(i int) {}
	)
	// 初始化dfs函数，用于遍历关系图
	dfs = func(i int) {
		// 无法完成提前结束
		if !result {
			return
		}
		// 当前路径上出现重复的节点，说明有循环关系，无法完成，直接结束
		if curPath[i] {
			result = false
			return
		}
		// 元素已经访问过了，不用重复走
		if visited[i] {
			return
		}
		// 标记当前节点为已访问过
		visited[i] = true
		// 当前节点走过当前路径
		curPath[i] = true
		for _, g := range graph[i] {
			dfs(g)
		}
		// 撤销当前节点的路径
		curPath[i] = false
	}
	// 遍历关系图
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	return result
}

// buildGraph 构建关系图
func buildGraph(size int, prerequisites [][]int) [][]int {
	graph := make([][]int, size)
	for _, pre := range prerequisites {
		from, to := pre[1], pre[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}
