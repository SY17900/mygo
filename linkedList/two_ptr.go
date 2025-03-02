package linkedlist

func Cross(ll1, ll2 *LinkedList) *ListNode {
	node1, node2 := ll1.Head, ll2.Head
	for node1 != node2 {
		if node1 != nil {
			node1 = node1.Next
		} else {
			node1 = ll2.Head
		}
		if node2 != nil {
			node2 = node2.Next
		} else {
			node2 = ll1.Head
		}
	}

	return node1
}

func WhereCircle(ll *LinkedList) *ListNode {
	slow, fast := ll.Head, ll.Head.Next.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = ll.Head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
