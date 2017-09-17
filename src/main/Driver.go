package main

import (
	"fmt"
	"mysort"
)

func main() {
	inputNumbers := []int{2, 3, 5, 8, 6, 9, 1}
	fmt.Println(mysort.MergeSort(inputNumbers))
}
