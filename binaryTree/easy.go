package binarytree

import "strconv"

func InvertTree(root *TreeNode) {
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		helper(root.Left)
		helper(root.Right)
	}

	helper(root)
}

func IfBalance(root *TreeNode) bool {
	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := helper(root.Left), helper(root.Right)
		if left == -1 || right == -1 {
			return -1
		}
		if left-right > 1 || right-left > 1 {
			return -1
		}
		return max(left, right) + 1
	}

	h := helper(root)
	return h != -1
}

func AllPath(root *TreeNode) (ans []string) {
	var helper func(*TreeNode, string)
	helper = func(root *TreeNode, s string) {
		if root.Left == nil && root.Right == nil {
			res := s + strconv.Itoa(root.Val)
			ans = append(ans, res)
			return
		}
		s = s + strconv.Itoa(root.Val) + "->"
		if root.Left != nil {
			helper(root.Left, s)
		}
		if root.Right != nil {
			helper(root.Right, s)
		}
	}

	helper(root, "")
	return
}
