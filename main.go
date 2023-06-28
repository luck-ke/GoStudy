package main

import (
	"GoStudy/algorithm/tree"
	"fmt"
)

func main() {
	//sli := []int{1, 23, 34}

	//sli := []int{5, 23, 34, 5, 12, 1, 23, 2, 23, 5, 762, 1, 23, 4}
	////testLen(sli[:7])
	//algorithnm.Quick(sli)
	//fmt.Println(sli)
	//
	//fmt.Println(dfs.Parenthesis(3))

	nodes := []tree.Node{
		{Id: 1, ParentId: 0, Value: "Node 1"},
		{Id: 2, ParentId: 1, Value: "Node 2"},
		{Id: 3, ParentId: 1, Value: "Node 3"},
		{Id: 4, ParentId: 2, Value: "Node 4"},
		{Id: 5, ParentId: 2, Value: "Node 5"},
		{Id: 6, ParentId: 3, Value: "Node 6"},
		{Id: 7, ParentId: 3, Value: "Node 7"},
		{Id: 8, ParentId: 4, Value: "Node 8"},
		{Id: 9, ParentId: 4, Value: "Node 9"},
	}
	for _, v := range *tree.FindLeafs(&nodes, 0) {
		fmt.Println(*v)
	}
}

func testLen(s []int) {
	testLen2(s[:3])
	fmt.Println(len(s))
}

func testLen2(s []int) {
	fmt.Println(len(s))
}
