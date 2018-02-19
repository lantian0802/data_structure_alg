package mysort

import (
	"fmt"
	"math"
	"strconv"
)

func RadixSort(arrayToSort []int, digit int) []int {
	//low to high digit
	for k := 1; k <= digit; k++ {
		//temp array to store the sort result inside digit
		tmpArray := make([]int, len(arrayToSort))
		//temp array for countingsort
		tmpCountingSortArray := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		//CountingSort
		for i := 0; i < len(arrayToSort); i++ {
			//split the specified digit from the element
			tmpSplitDigit := arrayToSort[i]/int(math.Pow(10, (float64)(k-1))) - (arrayToSort[i]/int(math.Pow(10, float64(k))))*10
			fmt.Println("tmpSplitDigit=" + strconv.Itoa(tmpSplitDigit) + " k=" + strconv.Itoa(k))
			tmpCountingSortArray[tmpSplitDigit] += 1
		}
		for m := 1; m < 10; m++ {
			tmpCountingSortArray[m] += tmpCountingSortArray[m-1]
		}
		//output the value to result
		for n := len(arrayToSort) - 1; n >= 0; n-- {
			tmpSplitDigit := arrayToSort[n]/int(math.Pow(10, float64(k-1))) - (arrayToSort[n]/int(math.Pow(10, float64(k))))*10
			tmpArray[tmpCountingSortArray[tmpSplitDigit]-1] = arrayToSort[n]
			tmpCountingSortArray[tmpSplitDigit] -= 1
		}
		//copy the digit-inside sort result to source array
		for p := 0; p < len(arrayToSort); p++ {
			arrayToSort[p] = tmpArray[p]
		}
	}
	return arrayToSort
}
