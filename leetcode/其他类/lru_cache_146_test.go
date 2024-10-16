package 其他类

import "testing"

// lru缓存，146，mid
// https://leetcode.cn/problems/lru-cache/description/
// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value
// 如果不存在，则向缓存中插入该组 key-value 。
// 如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
// 示例：
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4

func Test_LRUCache(t *testing.T) {

}

// LRUCache map+双向链表实现lru缓存
// 思路：
// 1.map是用来查找key的
// 2.使用双向链表是为了能快速模拟队列的更新操作，把最近使用到的数据放到top1的位置，剔除最不经常使用的数据
type LRUCache struct {
	size  int
	cpa   int
	cache map[int]*LinkedNode
	head  *LinkedNode
	tail  *LinkedNode
}

type LinkedNode struct {
	prev  *LinkedNode
	next  *LinkedNode
	key   int
	value int
}

func InitLinkNode(key int, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

// ConstructorLRUCache 构造lru对象
func ConstructorLRUCache(capacity int) LRUCache {
	lru := LRUCache{
		size:  0,
		cpa:   capacity,
		cache: make(map[int]*LinkedNode, capacity),
		head:  InitLinkNode(0, 0),
		tail:  InitLinkNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get 逻辑：
// 1.判断key在cache中是否存在
// 2.不存在返回-1
// 3.存在更新数据的使用频率
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		// 1.把当前节点移动到链表top1位置
		this.move2Top1(node)
		// 2.返回节点的value
		return node.value
	}
}

// Put 逻辑
// 1.判断key在cache中是否存在
// 2.不存在：
// 将key存储到cache中，然后新建一个node，将node直接放到链表top1位置，size++
// 同时判断此时cache的size是否超过了容量，超过容量，要剔除最老的数据，size--
// 3.存在，更新key对应的value数据，更新数据的使用率频率
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		// 新建一个node
		newNode := InitLinkNode(key, value)
		// 数据加入cache
		this.cache[key] = newNode
		// 新node直接放入top1位置
		this.add2Top1(newNode)
		this.size++
		// 判断容量是否超标
		if this.size > this.cpa {
			// 移除最后一个node
			removed := this.removeTail()
			// cache map中也要删除
			delete(this.cache, removed.key)
			// 容量+1
			this.size--
		}
	} else {
		// 更新key对应的value
		node.value = value
		// node移动到top1位置
		this.move2Top1(node)
	}
}

// move2Top1 移动node到top1位置
func (this *LRUCache) move2Top1(node *LinkedNode) {
	// 先把node在原有链表的关系移除
	this.removeNode(node)
	// 直接把node加入到top1为止
	this.add2Top1(node)
}

// add2Top1 将node直接放到top1位置
func (this *LRUCache) add2Top1(node *LinkedNode) {
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}

// removeTail 从尾部移除一个node
func (this *LRUCache) removeTail() *LinkedNode {
	last := this.tail.prev
	this.removeNode(last)
	return last
}

// removeNode 把node从当前位置移除
func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
