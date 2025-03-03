package stackandqueue

type CosplayStack struct {
	queue []int
}

func (q *CosplayStack) Push(val int) {
	q.queue = append(q.queue, val)
}

func (q *CosplayStack) Pop() (res int) {
	length := len(q.queue) - 1
	for i := range length {
		temp := q.queue[i]
		q.queue = q.queue[1:]
		q.queue = append(q.queue, temp)
	}

	res = q.queue[0]
	q.queue = q.queue[1:]
	return
}

func (q *CosplayStack) Top() int {
	temp := q.Pop()
	q.queue = append(q.queue, temp)
	return temp
}

func (q *CosplayStack) Empty() bool {
	return len(q.queue) == 0
}
