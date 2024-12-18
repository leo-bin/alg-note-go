package 其他类

import (
	"fmt"
	"reflect"
	"testing"
)

// 49，字母异位词分组，mid
// https://leetcode.cn/problems/group-anagrams/description/
// 给你一个字符串数组，请你将 字母异位词 组合在一起
// 可以按任意顺序返回结果列表
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词
//
// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 示例 2:
// 输入: strs = [""]
// 输出: [[""]]
//
// 示例 3:
// 输入: strs = ["a"]
// 输出: [["a"]]

func Test_GroupAnagrams(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Strs         []string
		ExpectResult [][]string
	}{
		0: {
			Name: "case 1",
			Strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			ExpectResult: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		1: {
			Name: "case 2",
			Strs: []string{""},
			ExpectResult: [][]string{
				{""},
			},
		},
		2: {
			Name: "case 3",
			Strs: []string{"a"},
			ExpectResult: [][]string{
				{"a"},
			},
		},
		3: {
			Name: "case 4",
			Strs: []string{"bdddddddddd", "bbbbbbbbbbc"},
			ExpectResult: [][]string{
				{"bdddddddddd"},
				{"bbbbbbbbbbc"},
			},
		},
	}

	for i, testCase := range testCaseList {
		strs := testCase.Strs
		expect := testCase.ExpectResult
		actual := groupAnagrams(strs)
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("test case %v,not passed,expect=%v,but actual=%v", i, expect, actual)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// groupAnagrams hash
// 关键是如何进行分组
// 考虑到如果是异位词，那么肯定拥有相同的字符，且他们的值在int数组中的位置和出现次数一定相同
// 那么可以根据此作为分组的依据
func groupAnagrams(strs []string) [][]string {
	// build hash record
	resultMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		record := make([]int, 26)
		for j := 0; j < len(str); j++ {
			record[str[j]-'a']++
		}
		key := ""
		for j := 0; j < len(record); j++ {
			if record[j] > 0 {
				key += string(byte(j+'a')) + fmt.Sprintf("%v", record[j])
			}
		}
		if _, ok := resultMap[key]; ok {
			resultMap[key] = append(resultMap[key], str)
		} else {
			resultMap[key] = []string{str}
		}
	}
	// build result
	result := make([][]string, 0)
	for _, strList := range resultMap {
		result = append(result, strList)
	}
	return result
}
