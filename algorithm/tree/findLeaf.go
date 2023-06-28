package tree

import "fmt"

func FindLeafs(nodes *[]Node, id int) *[]*Node {
	leaf := make([]*Node, 0)
	idMap := make(map[int]*Node)
	pidMap := make(map[int]struct{})
	for _, v := range *nodes {
		node := v
		idMap[node.Id] = &node
		pidMap[node.ParentId] = struct{}{}
	}
	fmt.Println(idMap)
	fmt.Println(pidMap)
	buds := []*Node{}
	for _, v := range *nodes {
		_, ok := pidMap[v.Id]
		if ok {
			n := v
			buds = append(buds, &n)
		}
	}
	fmt.Println(len(buds))
	for _, v := range buds {
		leaf = append(leaf, getLeaf(nodes, &idMap, &pidMap, v))
	}
	return &leaf
}

func getLeaf(nodes *[]Node, idMap *map[int]*Node, pidMap *map[int]struct{}, node *Node) *Node {
	_, ok := (*pidMap)[(*node).Id]
	if ok {
		//return getLeaf(idMap, pidMap, (*idMap)[(*node).Id])
		for _, n := range *nodes {
			if n.ParentId == (*node).Id {
				return getLeaf(nodes, idMap, pidMap, &n)
			}
		}
	}
	return (*idMap)[(*node).Id]
}

type Node struct {
	Id       int
	ParentId int
	Value    string
}
