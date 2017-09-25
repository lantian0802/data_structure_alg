package mysort

func binarySearch(input []int,item int,low int,high int) int {
	if high <= low {
		if item > input[low] {
			return low + 1
		} else {
			return low
		}
	}
	middle := (high + low)/2
	if item == input[middle] {
		return middle + 1
	}

	if item > input[middle] {
		return binarySearch(input,item,middle+1,high)
	} else {
		return binarySearch(input, item, low, middle - 1)
	}
}

func BinaryInsertSort(input []int) []int {

	for i:=1; i < len(input); i++ {
		j := i-1
		index := binarySearch(input,input[i],0,j)
		for j >= index {
			input[j+1] = input[j]
			j--
		}
		input[index+1] = input[i]
	}
	return input
}


