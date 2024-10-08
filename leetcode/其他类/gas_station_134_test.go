package 其他类

import "testing"

// 加油站，134，mid
// https://leetcode.cn/problems/gas-station/description/
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升
// 你从其中的一个加油站出发，开始时油箱为空。
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号
// 否则返回 -1 。如果存在解，则 保证 它是 唯一 的
//
// 示例 1:
// 输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// 输出: 3
// 解释:
// 从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
// 开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
// 开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
// 开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
// 开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
// 开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
// 因此，3 可为起始索引。
//
// 示例 2:
// 输入: gas = [2,3,4], cost = [3,4,3]
// 输出: -1
// 解释:
// 你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
// 我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
// 开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
// 开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
// 你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
// 因此，无论怎样，你都不可能绕环路行驶一周

func Test_CanCompleteCircuit(t *testing.T) {

}

// canCompleteCircuit 贪心
// 思路：
// 1.模拟从每一个加油站开始往后移动，记录累计加油量和耗油量
// 2.如果经过某一加油站时，此时的累计加油量比耗油量要小，说明后面的路无法走下去，从下一个加油站从头开始计算
// 3.假设从x出发，经过y，最多能到达z，那么从y开始也最多只能到达z
// 4.因为从x出发是有初始油量的，而从y开始，此时没有初始油量，更加没法到达z了
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for start := 0; start < n; {
		sumGas, sumCost, cnt := 0, 0, 0
		for cnt < n {
			i := (start + cnt) % n
			sumGas += gas[i]
			sumCost += cost[i]
			// 油量不够，终止
			if sumCost > sumGas {
				break
			}
			cnt++
		}
		// 走完全部，返回开始下标
		if cnt == n {
			return start
		} else {
			// 没走完，从cnt+1开始往后找
			start += cnt + 1
		}
	}
	return -1
}
