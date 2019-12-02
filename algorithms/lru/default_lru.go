package lru

import (
	"fmt"
	"github.com/DestinyWang/go-widget/data_structs/linked_list"
)

const DefaultSize = 1000

// 默认 LRU 算法, 底层由单链表实现
type DefaultLru struct {
	cap  int
	size int
	list *linked_list.LinkedList
}

func NewDefaultLru(size int) (lru *DefaultLru) {
	return &DefaultLru{
		cap:  size,
		size: 0,
		list: &linked_list.LinkedList{},
	}
}

func (lru *DefaultLru) Push(data string) {
	if lru == nil {
		lru = NewDefaultLru(DefaultSize)
	}
	pos := lru.list.Pos(data)
	if pos >= 0 {
		// 如果存在
		lru.list.Del(pos)
		lru.list.Push(data)
	} else {
		// 如果不存在, 删除最后一个节点, 再插入当前节点
		if lru.size >= lru.cap {
			lru.list.Pop()
		}
		lru.list.Push(data)
		lru.size++
	}
}

func (lru *DefaultLru) String() string {
	return fmt.Sprintf("lru cap=[%d], lru size=[%d]\nlru list=[%s]", lru.cap, lru.size, lru.list.String())
}
