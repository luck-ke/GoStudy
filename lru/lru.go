package lru

import (
	"container/list"
)

type LRUNode struct {
	Caps int
	Key  map[int]list.Element
	list *list.List
}

type hash struct {
	Key, Value int
}

type lRU struct {
}

func (lRU) New(caps int) *LRUNode {
	return &LRUNode{
		Caps: caps,
		Key:  make(map[int]list.Element),
		list: list.New(),
	}
}

func (node LRUNode) Get(key int) int {
	value, ok := node.Key[key]
	if ok {
		node.list.MoveToFront(&value)
		return value.Value.(hash).Value
	}
	return -1
}

func (node LRUNode) Set(key, value int) {
	elem, ok := node.Key[key]
	if ok {
		elem.Value = hash{
			Key:   key,
			Value: value,
		}
		node.list.MoveToFront(&elem)
	} else {
		if node.Caps < node.list.Len() {
			delete(node.Key, node.list.Back().Value.(hash).Value)
			node.list.Remove(node.list.Back())
		}
		node.list.PushFront(hash{key, value})
		node.Key[key] = *node.list.Front()
	}

}
