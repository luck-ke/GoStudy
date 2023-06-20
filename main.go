package main

import (
	"GoStudy/algorithm/dfs"
	algorithnm "GoStudy/algorithm/sort"
	"fmt"
)

func main() {
	//sli := []int{1, 23, 34}

	sli := []int{5, 23, 34, 5, 12, 1, 23, 2, 23, 5, 762, 1, 23, 4}
	//testLen(sli[:7])
	algorithnm.Quick(sli)
	fmt.Println(sli)

	fmt.Println(dfs.Parenthesis(3))
}

func testLen(s []int) {
	testLen2(s[:3])
	fmt.Println(len(s))
}

func testLen2(s []int) {
	fmt.Println(len(s))
}
