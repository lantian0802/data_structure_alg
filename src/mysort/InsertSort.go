package mysort

//插入排序
//时间复杂度 O(n) * O(n) = O(n^2)
func InsertSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		for j := i; j > 0; j-- {
			if input[j] < input[j-1] {
				input[j-1],input[j] = input[j],input[j-1]
			}

		}
	}
	return input
}
