package search

type Node struct {
	value int
	left  *Node
	right *Node
}

//用O(N)时间构造平衡二查搜索树
func BuildBinarySearchTree(nums []int, start int, end int) *Node {
	if end < start {
		return nil
	}
	middle := (end - start) / 2
	root := Node{middle, nil, nil}
	root.left = BuildBinarySearchTree(nums, start, middle-1)
	root.left = BuildBinarySearchTree(nums, middle+1, end)
	return &root
}
