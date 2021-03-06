package mysort
import "fmt"

// 时间复杂度 O(n) * O(lgn) = O(nlgn)
// 空间复杂度 O(n)
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
		item := input[i]
		fmt.Printf("index:%v, input:%v \n",index,input)
		for j >= index {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = item
		fmt.Printf("input[j+1]:%v,input[i]:%v,j=%v,i=%v \n",input[j+1],input[i],j,i)
	}
	return input
}


