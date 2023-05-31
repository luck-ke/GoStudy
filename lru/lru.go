package lru

import "container/list"

type LRUNode struct {
	Caps int
	Key  map[int]list.Element
	list *list.List
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

func (lRU) Get() {

}
