package main

import (
	"fmt"
	//"leetcode"
	//"strings"
	"mysort"
)

func main() {
	var numList = []int{1, 3, 2, 7, 6, 4}
	//0 1 1 1 1 0 1 1
	fmt.Println(mysort.CountSort(numList))
}
