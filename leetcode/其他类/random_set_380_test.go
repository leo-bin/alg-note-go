package 其他类

import (
	"math/rand"
	"testing"
)

// O(1)时间插入、删除以及获取随机元素，380，mid
// https://leetcode.cn/problems/insert-delete-getrandom-o1/description/
// 实现RandomizedSet 类：
// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1)
//
// 示例：
// 输入
// ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// 输出
// [null, true, false, true, 2, true, false, 2]
// 解释
// RandomizedSet randomizedSet = new RandomizedSet();
// randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
// randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
// randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
// randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
// randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
// randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
// randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2

func Test_RandomSet(t *testing.T) {

}

type RandomizedSet struct {
	nums        []int
	valIndexMap map[int]int
}

func ConstructorRandomSet() RandomizedSet {
	return RandomizedSet{
		nums:        make([]int, 0),
		valIndexMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valIndexMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.valIndexMap[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.valIndexMap[val]; !ok {
		return false
	}
	index := this.valIndexMap[val]
	last := this.nums[len(this.nums)-1]
	// 更新indexMap中最后一个元素的index
	this.valIndexMap[last] = index
	// 交换待删除元素和最后一个元素
	this.nums[index], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[index]
	// 删除此时的最后一个元素
	this.nums = this.nums[:len(this.nums)-1]
	// 删除indexMap中的记录
	delete(this.valIndexMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
