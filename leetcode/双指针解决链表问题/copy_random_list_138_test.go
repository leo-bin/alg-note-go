package 双指针解决链表问题

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 复制链表，138，mid
// 复制带随随机节点的链表
// https://leetcode.cn/problems/copy-list-with-random-pointer/
// 思路：
// 1.二次遍历+hash表即可
// 2.注意新链表的random应该也指向random
func copyRandomList(head *Node) *Node {
	// base case
	if head == nil {
		return nil
	}

	dummy := &Node{}
	d := dummy
	p := head
	hash := make(map[*Node]*Node)

	// 1.第一次遍历，构建新节点以及记录新老节点的关系
	for p != nil {
		d.Next = &Node{
			Val: p.Val,
		}
		hash[p] = d.Next

		// p和d往后走
		p = p.Next
		d = d.Next
	}

	// 2.第二次遍历，通过新老关系找新节点的random应该指向新链表的哪一个节点
	d = dummy.Next
	p = head
	for p != nil {
		if p.Random != nil {
			if newNode, ok := hash[p.Random]; ok {
				d.Random = newNode
			}
		}
		d = d.Next
		p = p.Next
	}

	return dummy.Next
}
