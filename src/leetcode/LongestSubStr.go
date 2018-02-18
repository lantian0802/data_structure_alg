package leetcode

import (
	"fmt"
)

func LengthOfLongestSubstring(s string) int {
	hashmap := map[byte]int{}
	max := 0
	for i := range s {
		a, ok := hashmap[s[i]]
		fmt.Println(a)
		if !ok {
			hashmap[s[i]] = i
			if len(hashmap) > max {
				max = len(hashmap)
			}
		} else {
			// remove repeated
			oldI := hashmap[s[i]]
			hashmap[s[i]] = i

			for key, value := range hashmap {
				if value < oldI {
					delete(hashmap, key)
				}
			}
		}
	}
	return max
}
