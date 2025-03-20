package binarytree

func CommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	var helper func(*TreeNode) *TreeNode
	helper = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root == p || root == q {
			return root
		}
		left, right := helper(root.Left), helper(root.Right)
		if left != nil && right != nil {
			return root
		}
		if left != nil {
			return left
		}
		if right != nil {
			return right
		}
		return nil
	}

	return helper(root)
}
