package mysort

// SelectionSort
// 时间复杂度 O(n)* O(n) = O(n^2)
// 空间复杂度 O(n)
//外层循环只是用来控制次数
//内层循环保证比较和交换
func SelectionSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		temp := input[i]
		pos := i
		for j := i + 1; j < len(input); j++ {
			if temp > input[j] {
				pos = j
			}
		}
		input[i], input[pos] = input[pos], input[i]
	}
	return input
}
