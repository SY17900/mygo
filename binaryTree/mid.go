package binarytree

func MergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	q := []*TreeNode{t1, t2}
	for len(q) != 0 {
		cur1, cur2 := q[0], q[1]
		q = q[2:]
		cur1.Val += cur2.Val
		if cur1.Left != nil && cur2.Left != nil {
			q = append(q, cur1.Left)
			q = append(q, cur2.Left)
		}
		if cur1.Right != nil && cur2.Right != nil {
			q = append(q, cur1.Right)
			q = append(q, cur2.Right)
		}
		if cur1.Left == nil {
			cur1.Left = cur2.Left
		}
		if cur1.Right != nil {
			cur1.Right = cur2.Right
		}
	}

	return t1
}
