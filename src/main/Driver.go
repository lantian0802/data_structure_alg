package main

import (
	"fmt"
	//"leetcode"
	//"strings"
	//"mysort"
	"search"
)

func main() {
	var numList = []int{1, 2, 3, 4, 5, 6}
	//var nums1 = []int{1, 2}
	//var nums2 = []int{2}
	//0 1 1 1 1 0 1 1
	fmt.Println(search.BuildBinarySearchTree(numList, 0, 5))
	//leetcode.FindMedianSortedArrays(nums1, nums2)
}
