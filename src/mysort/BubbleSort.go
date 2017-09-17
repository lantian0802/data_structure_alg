package mysort

// BubbleSort
// 时间复杂度 O(n)* O(n) = O(n^2)
// 空间复杂度 O(n)
//外层循环只是用来控制次数
//内层循环保证比较和交换
func BubbleSort(input []int) []int {
	for i := len(input) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

func BubbleSortV2(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
