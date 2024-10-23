package 其他类

import "testing"

// 跳跃游戏ii，45，mid
// https://leetcode.cn/problems/jump-game-ii/description/
// 给定一个长度为 n 的 0 索引整数数组 nums
// 初始位置为 nums[0]
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度
// 换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
// 0 <= j <= nums[i] ，i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数
// 生成的测试用例可以到达 nums[n - 1]
//
// 示例 1:
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
// 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置
//
// 示例 2:
// 输入: nums = [2,3,0,1,4]
// 输出: 2

func Test_JumpGameII(t *testing.T) {
	testCaseList := []struct {
		Name         string
		Nums         []int
		ExpectResult int
	}{
		0: {
			Name:         "case 1",
			Nums:         []int{2, 3, 1},
			ExpectResult: 1,
		},
		1: {
			Name:         "case 2",
			Nums:         []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3},
			ExpectResult: 2,
		},
	}

	for i, testCase := range testCaseList {
		nums := testCase.Nums
		expectResult := testCase.ExpectResult
		actualResult := jumpGameII(nums)
		if actualResult != expectResult {
			t.Errorf("test case %v,not passed,expectResult=%v,but actualResult=%v", i, expectResult, actualResult)
		} else {
			t.Logf("test case %v,passed", i)
		}
	}
}

// jumpGameII 贪心
// 思路：
// 1.依旧是动态的寻找最新的最远能跳到的距离，假设是maxJump
// 2.我们在寻找maxJump过程中，一定会找到一个位置能够更新maxJump，那么最优跳跃点一定会在i-maxJump之间产生
// 3.因为最优跳跃点只有1个，所以不需要找到最优跳跃点在哪个具体的位置，我们只需要知道什么时候离开了maxJump即可
// 4.使用一个临时变量，lastMaxJump来记录上一次的最原远距离
// 5.当经过了lastMaxJump时，就说明经过了一个最优跳跃区间，step+1即可
func jumpGameII(nums []int) int {
	maxJump, lastMaxJump, minStep := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxJump = maxJumpII(maxJump, i+nums[i])
		if i == lastMaxJump {
			lastMaxJump = maxJump
			minStep++
		}
	}
	return minStep
}

func maxJumpII(x, y int) int {
	if x > y {
		return x
	}
	return y
}
