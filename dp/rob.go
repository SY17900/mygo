package dp

import "slices"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Rob(root *TreeNode) int {
	var robHelper func(*TreeNode) []int // 0: no, 1: yes
	robHelper = func(rt *TreeNode) []int {
		if rt == nil {
			return []int{0, 0}
		}
		leftRes, rightRes := robHelper(rt.Left), robHelper(rt.Right)
		return []int{
			slices.Max(leftRes) + slices.Max(rightRes),
			rt.Val + leftRes[0] + rightRes[0],
		}
	}

	return slices.Max(robHelper(root))
}
