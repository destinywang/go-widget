package cn

//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。 
//写入数据 put(key, value) - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写
//入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。 
//
// 
//
// 进阶: 
//
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？ 
//
// 
//
// 示例: 
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );	nil
// cache.put(1, 1);										nil
// cache.put(2, 2);										nil
// cache.get(1);       // 返回  1						1
// cache.put(3, 3);    // 该操作会使得密钥 2 作废		nil
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得密钥 1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4
// 
// Related Topics 设计

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	cap  int
	len  int
	head *Node
	tail *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	var p = this.head
	for ; p != nil; p = p.next {
		if p.key == key {
			if p != this.head {
				// remove p to head
				this.head.prev = p
				if p.next != nil {
					p.next = this.head
				}
				this.head = p
				p.prev.next = p.next
				p.next.prev = p.prev
			}
			return p.value
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.len >= this.cap {
		// 删除当前 tail 节点
		this.tail = this.tail.prev
		this.tail.next = nil
	}
	n := &Node{
		key:   key,
		value: value,
	}
	if this.head == nil {
		// insert as head
		this.head = n
		this.tail = n
	} else {
		// insert to head
		this.head.prev = n
		n.next = this.head
		this.head = n
	}
	this.len++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
