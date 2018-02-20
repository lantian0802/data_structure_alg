package mysort

import (
	"fmt"
	"strconv"
)

const BUCKET_NUM = 5
const INTERVAL = 10

func GetBucketIndex(value int) int {
	checkValue := value / INTERVAL
	switch checkValue {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	default:
		return 4
	}
}

func BucketSort(numList []int) []int {
	if len(numList) == 0 {
		return numList
	}
	bucketMap := map[int][]int{}
	//数据入桶
	for i := 0; i < len(numList); i++ {
		fmt.Println("currentNum=" + strconv.Itoa(numList[i]))
		subList, ok := bucketMap[GetBucketIndex(numList[i])]
		if ok {
			fmt.Println("step0")
			subList = append(subList, numList[i])
			bucketMap[GetBucketIndex(numList[i])] = subList
			fmt.Println("subList.size=" + strconv.Itoa(len(subList)))
		} else {
			fmt.Println("step1")
			subList = make([]int, 1)
			subList[0] = numList[i]
			bucketMap[GetBucketIndex(numList[i])] = subList
			fmt.Println("subList.size=" + strconv.Itoa(len(subList)))
		}
	}
	//每个桶的数据进行排序
	for j := 0; j < BUCKET_NUM; j++ {
		bucketMap[j] = InsertSort(bucketMap[j])
	}
	//将所有桶的数据输出到原有的数组中
	index := 0
	for k := 0; k < BUCKET_NUM; k++ {
		subList := bucketMap[k]
		for n := 0; n < len(subList); n++ {
			numList[index] = subList[n]
			index += 1
		}
	}
	return numList
}
