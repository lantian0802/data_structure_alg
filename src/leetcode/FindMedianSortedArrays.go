package leetcode

import (
	"fmt"
	"strconv"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var totalNums = make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0
	k := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			totalNums[k] = nums1[i]
			i++
		} else {
			totalNums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) {
		totalNums[k] = nums1[i]
		k++
		i++
	}

	for j < len(nums2) {
		totalNums[k] = nums2[j]
		k++
		j++

	}

	median := 0.0
	fmt.Println("totoalNums=" + strconv.Itoa(len(totalNums)))
	if len(totalNums)%2 == 0 {
		index1 := len(totalNums)/2 - 1
		index2 := index1 + 1
		fmt.Println("index1=" + strconv.Itoa(index1) + " index2=" + strconv.Itoa(index2))
		median = (float64(totalNums[index1]) + float64(totalNums[index2])) / 2.0
	} else {
		index1 := len(totalNums) / 2
		median = float64(totalNums[index1])
	}

	for m := 0; m < len(totalNums); m++ {
		fmt.Println(totalNums[m])
	}

	fmt.Println("median=" + strconv.FormatFloat(median, 'f', 1, 64))
	return median
}
