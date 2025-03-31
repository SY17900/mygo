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
		if cur1.Right == nil {
			cur1.Right = cur2.Right
		}
	}

	return t1
}

func BuildTree(inorder []int, postorder []int) *TreeNode {
	mp := make(map[int]int)
	for i, num := range inorder {
		mp[num] = i
	}
	var helper func(int, int, int) *TreeNode
	helper = func(rootIndexPost, leftIn, rightIn int) *TreeNode {
		if leftIn > rightIn {
			return nil
		}
		if leftIn == rightIn {
			return &TreeNode{Val: inorder[leftIn]}
		}
		rootNum := postorder[rootIndexPost]
		rootIndexIn := mp[rootNum]
		root := &TreeNode{Val: rootNum}
		root.Right = helper(rootIndexPost-1, rootIndexIn+1, rightIn)
		root.Left = helper(rootIndexPost-(rightIn-rootIndexIn)-1, leftIn, rootIndexIn-1)

		return root
	}

	return helper(len(postorder)-1, 0, len(inorder)-1)
}
