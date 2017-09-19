package mysort
import "fmt"
// BubbleSort
// 时间复杂度 O(n)* O(lgn) = O(nlgn)
//外层循环只是用来控制次数
//内层循环保证比较和交换
func QuickSort(input [] int,left int,right int) []int {
	if left < right {
		low,high := left,right
		key := input[low]
		for low < high {
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
		fmt.Println("====result====")
		fmt.Println(input)
		fmt.Println("========")
		fmt.Println("====left sub====")
		fmt.Println(input[0:low])
		fmt.Println("========")
		fmt.Println("====right sub====")
		fmt.Println(input[low+1:right+1])
		fmt.Println("========")
		QuickSort(input,left,low-1)
		QuickSort(input,low+1,right)
	}

	return input

}
