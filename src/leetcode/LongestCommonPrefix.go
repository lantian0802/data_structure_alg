package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	tempCount := len(strs[0])
	tempIndex := 0
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < tempCount {
			tempCount = len(strs[i])
			tempIndex = i
			break
		}
	}
	tempStr := strs[tempIndex]
	fmt.Println("tempStr=" + tempStr + " tempIndex=" + strconv.Itoa(tempIndex))
	lonestCommonPrefix := ""
	for j := 0; j <= len(tempStr); j++ {
		var flag = true
		var prefix = tempStr[0:j]
		fmt.Println(prefix)
		fmt.Println("j=" + strconv.Itoa(j))
		for k := 0; k < len(strs); k++ {
			if k == tempIndex {
				continue
			}
			fmt.Println("strs[k]=" + strs[k] + " prefix=" + prefix + " result=" + strconv.Itoa(strings.Index(strs[k], prefix)))
			if strings.Index(strs[k], prefix) == 0 {
				fmt.Println("step0")
				continue
			} else {
				fmt.Println("step1")
				flag = false
				break
			}
		}
		if flag {
			lonestCommonPrefix = prefix
		}
	}
	return lonestCommonPrefix
}
