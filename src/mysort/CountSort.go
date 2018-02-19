package mysort

import (
	"fmt"
	"strconv"
)

// Time Complexty O(N)
// Space Complexty O(N)
//计数排序
func CountSort(numList []int) []int {
	maxValue := 0
	for i := 0; i < len(numList); i++ {
		if maxValue < numList[i] {
			maxValue = numList[i]
		}
	}
	//初始化为0
	var indexArray = make([]int, maxValue+1)
	//统计元素个数
	for k := 0; k < len(numList); k++ {
		indexArray[numList[k]] = indexArray[numList[k]] + 1
	}
	//统计小于当前元素的个数
	for m := 1; m < len(indexArray); m++ {
		indexArray[m] = indexArray[m] + indexArray[m-1]
	}
	//各归各位
	result := make([]int, len(numList)+1)
	for n := len(numList) - 1; n >= 0; n-- {
		fmt.Println(strconv.Itoa(numList[n]))
		fmt.Println("indexValue = " + strconv.Itoa(indexArray[numList[n]]))
		result[indexArray[numList[n]]] = numList[n]
		indexArray[numList[n]] = indexArray[numList[n]] - 1
	}
	return result[len(result)-len(numList) : len(result)]
}
