package mysort

func HeapSort(numList []int) []int {
	//建堆
	for i := len(numList)/2 - 1; i >= 0; i-- {
		AdjustHeap(numList, i, len(numList)-1)
	}

	//堆排序
	for j := len(numList) - 1; j >= 0; j-- {
		temp := numList[0]
		numList[0] = numList[j]
		numList[j] = temp
		AdjustHeap(numList, 0, j-1)
	}
	return numList
}

//调整堆
func AdjustHeap(numList []int, head int, length int) {
	temp := numList[head]
	index := 0

	for index = 2 * head; index < length; index *= 2 {
		if index < length && numList[index] < numList[index+1] {
			index++
		}

		if temp >= numList[index] {
			break
		}

		numList[head] = numList[index]
		head = index
	}
	numList[head] = temp
}
