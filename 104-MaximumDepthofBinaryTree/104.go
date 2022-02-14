package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Code starts here
func maxDepth(root *TreeNode) int {
	var outerqueue []*TreeNode
	var innerqueue []*TreeNode
	count := 0

	if root == nil {
		return count
	}

	if root != nil {
		count++
		if root.Left != nil {
			outerqueue = append(outerqueue, root.Left)
		}
		if root.Right != nil {
			outerqueue = append(outerqueue, root.Right)
		}
	}

	innerqueue, count = NextSet(outerqueue, count)

	for len(innerqueue) != 0 {
		outerqueue = innerqueue
		innerqueue, count = NextSet(outerqueue, count)
	}
	return count
}

func NextSet(inner []*TreeNode, count int) ([]*TreeNode, int) {
	var outer []*TreeNode
	if len(inner) == 0 {
		return outer, count
	}
	for _, node := range inner {
		if node.Left != nil {
			outer = append(outer, node.Left)
		}
		if node.Right != nil {
			outer = append(outer, node.Right)
		}
	}
	return outer, count + 1

}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// 104. Maximum Depth of Binary Tree
