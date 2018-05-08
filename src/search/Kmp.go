package search

import (
	"fmt"
	"strconv"
)

//KMP 算法核心
func KmpMatcher(target string, pattern string) int {
	n := len(target)
	m := len(pattern)
	next := ComputePrefixFunction(pattern)
	q := 0
	for i := 1; i < n; i++ {
		for q > 0 && pattern[q+1] != target[i] {
			q = next[q]
		}
		if pattern[q+1] == target[i] {
			q = q + 1
		}
		if q == m {
			fmt.Println("Pattern occurs with shift:" + (i - m))
		}
		q = next[q]
	}

}

// 1 2 3 4 5 6 7
// 0 1 2 3 4 5 6
// a b a b a c a
//计算最大前缀
func ComputePrefixFunction(pattern string) []int {
	m := len(pattern)
	next := make([]int, len(pattern))
	if m == 1 {
		return next
	}
	if pattern[0] == pattern[1] {
		next[1] = 1
	} else {
		next[1] = 0
	}
	for q := 2; q < m; i++ {
		if pattern[q] == pattern[next[q-1]] {
			next[q] = next[q-1] + 1
		} else {
			if pattern[i] == pattern[0] {
				next[q] = 1
			}
		}
	}
	return next
}
