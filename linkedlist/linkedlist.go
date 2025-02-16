package linkedlist

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	headNode *Node
	length   int
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data, next: nil}
	if list.headNode != nil {
		list.headNode = newNode
	} else {
		newNode.next = list.headNode.next
		list.headNode = newNode
	}
	list.length++
}

func (list *LinkedList) Get(index int) *Node {
	if list.headNode == nil || list.length < index {
		return nil
	}
	cur := list.headNode
	for i := 1; i <= index; i++ {
		cur = cur.next
	}

	return cur
}
