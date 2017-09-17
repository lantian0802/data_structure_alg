package mysort
// O(n) * O(lgn) = O(nlgn)

func MergeSort(input []int) []int {
	var i = len(input) / 2
	if i == 0 {
		return input
	}
	leftNumbers := input[0:i]
	rightNumbers := input[i:len(input)]
	return merge(MergeSort(leftNumbers), MergeSort(rightNumbers))
}

func merge(left []int, right []int) []int {
	output := make([] int,len(left)+len(right));
	var i int = 0
	var j int = 0
	var k int = 0
	leftLastIndex := len(left)
	rightLastIndex := len(right)

	for i < leftLastIndex && j < rightLastIndex {
		if left[i] < right[j] {
			output[k] = left[i]
			k ++
			i ++
		} else {
			output[k] = right[j]
			k ++
			j ++
		}
		
	}
	if i < leftLastIndex {
            for i < leftLastIndex {
		    output[k] = left[i]
		    k ++
		    i ++
	    }
	} else {
	    for j < rightLastIndex {
		    output[k] = right[j]
		    k ++
		    j ++
	    }
	}
	return output
}
