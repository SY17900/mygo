package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

func (ll *LinkedList) RemoveAll(val int) *LinkedList {
	dummy := &ListNode{}
	dummy.Next = ll.Head
	cur := dummy

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			ll.Size--
		} else {
			cur = cur.Next
		}
	}

	ll.Head = dummy.Next
	return ll
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || index >= ll.Size {
		return -1
	}
	cur := ll.Head
	for ; index > 0; index-- {
		cur = cur.Next
	}

	return cur.Val
}

func (ll *LinkedList) AddAtHead(val int) {
	ll.AddAtIndex(0, val)
}

func (ll *LinkedList) AddAtTail(val int) {
	ll.AddAtIndex(ll.Size, val)
}

func (ll *LinkedList) AddAtIndex(index, val int) int {
	if index < 0 || index > ll.Size {
		return -1
	}
	pre := &ListNode{Next: ll.Head}
	for ; index > 0; index-- {
		pre = pre.Next
	}
	pre.Next = &ListNode{Val: val, Next: pre.Next}
	ll.Size++

	return 0
}

func (ll *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= ll.Size {
		return
	}

	pre := &ListNode{Next: ll.Head}
	for ; index > 0; index-- {
		pre = pre.Next
	}
	temp := pre.Next
	pre.Next = temp.Next
	ll.Size--
}

func (ll *LinkedList) PrintOut() {
	cur := ll.Head
	for ; cur != nil; cur = cur.Next {
		print(cur.Val)
		if cur.Next != nil {
			print(" ")
		}
	}
}

func (ll *LinkedList) Reverse() {
	if ll.Size <= 1 {
		return
	}
	var pre *ListNode
	cur := ll.Head
	for cur != nil {
		temp := cur.Next
		cur.Next, pre = pre, cur
		cur = temp
	}
	ll.Head = pre
}

func (ll *LinkedList) Exchange() {
	dummy := &ListNode{Next: ll.Head}
	pre := dummy
	cur := pre.Next
	for cur != nil {
		newCur := cur.Next.Next
		cur.Next.Next = cur
		pre.Next = cur.Next
		cur.Next = newCur
		pre = cur
		cur = newCur
	}
	ll.Head = dummy.Next
}

func (ll *LinkedList) RemoveFromEnd(index int) *ListNode {
	slow, fast := ll.Head, ll.Head

	for ; index > 0; index-- {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
