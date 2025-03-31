package binarytree

func TrimBST0(root *TreeNode, low, high int) *TreeNode {
	var helper func(*TreeNode) *TreeNode
	helper = func(r *TreeNode) *TreeNode {
		if r == nil {
			return nil
		}
		if r.Val < low {
			return helper(r.Right)
		}
		if r.Val > high {
			return helper(r.Left)
		}
		r.Left = helper(r.Left)
		r.Right = helper(r.Right)

		return r
	}

	return helper(root)
}

func TrimBST1(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}

	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}

	return root
}
