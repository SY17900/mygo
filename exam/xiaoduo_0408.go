package exam

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder(root *TreeNode) (ans [][]int) {
    if root == nil {
		return
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	rec := 0
	
	for len(q) != 0 {
		if rec%2 == 0 {
			// 从左往右
			length := len(q)
			temp := make([]int, 0, length)
			for range length {
				cur := q[0]
				q = q[1:]
				temp = append(temp, cur.Val)
				if cur.Left != nil {
					q = append(q, cur.Left)
				}
				if cur.Right != nil {
					q = append(q, cur.Right)
				}
			}
			ans = append(ans, temp)
		} else {
			// 从右往左
			length := len(q)
			temp := make([]int, 0, length)
			for range length {
				cur := q[len(q)-1]
				q = q[:len(q)-1]
				temp = append(temp, cur.Val)
				if cur.Right != nil {
					q = append([]*TreeNode{cur.Right}, q...)
				}
				if cur.Left != nil {
					q = append([]*TreeNode{cur.Left}, q...)
				}
			}
			ans = append(ans, temp)
		}
        rec++
	}

	return
}
