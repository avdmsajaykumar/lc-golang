package internal

func maxd(root *TreeNode, l int) int {
	if root == nil {
		return l
	}
	left := maxd(root.Left, l+1)
	right := maxd(root.Right, l+1)
	if left > right {
		return left
	}
	return right
}
func MaxDepth(root *TreeNode) int {
	return maxd(root, 0)
}
