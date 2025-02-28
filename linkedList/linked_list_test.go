package linkedlist

import "testing"

func TestReverse(t *testing.T) {
	ll := LinkedList{
		Head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		Size: 4,
	}
	println(ll.RemoveFromEnd(1).Val)
}
