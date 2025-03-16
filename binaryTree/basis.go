package binarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (t *TreeNode) PreOrder() (res []int) {
	if t == nil {
		return
	}

	stack := []*TreeNode{}
	stack = append(stack, t)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return
}

func (t *TreeNode) InOrder() (res []int) {
	if t == nil {
		return
	}

	stack := []*TreeNode{}
	cur := t
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		stack = append(stack, cur.Right)
	}

	return
}

func (t *TreeNode) PostOrder() (res []int) {
	if t == nil {
		return
	}

	stack := []*TreeNode{}
	cur := t
	var pre *TreeNode
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		if cur.Right == nil || cur.Right == pre {
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			pre = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}

	return
}
