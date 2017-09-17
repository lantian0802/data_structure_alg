// BubbleSort
// 时间复杂度 O(n)* O(n) = O(n^2)
// 空间复杂度 O(n)
package main

import (
	"fmt"
)

func main() {
	inputNumbers := []int{2, 3, 5, 8, 6, 9, 1}
	fmt.Println(bubbleSortV2(inputNumbers))
}

//外层循环只是用来控制次数
//内层循环保证比较和交换
func bubbleSort(input []int) []int {
	for i := len(input) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

func bubbleSortV2(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
