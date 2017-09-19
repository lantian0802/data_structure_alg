package mysort

// BubbleSort
// 时间复杂度 O(n)* O(n) = O(n^2)
// 空间复杂度 O(n)
//外层循环只是用来控制次数
//内层循环保证比较和交换
func QuickSort(input [] int,left int,right int) []int {
	low,high := left,right
	key := input[low]
	for left < right {
		for input[high] > key && low < high {
			high --
		}
		input[low] = input[high]

		for input[low] < key && low < high {
			low ++
		}
		input[high] = input[low]
	}
	input[low] = key
	QuickSort(input[0:low],left,low-1)
	QuickSort(input[low+1:right],low+1,right)

	return input

}
