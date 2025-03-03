package stackandqueue

type CosplayQueue struct {
	stack1 []int
	stack2 []int
}

func (s *CosplayQueue) Push(val int) {
	s.stack1 = append(s.stack1, val)
}

func (s *CosplayQueue) Peek(val int) int {
	temp := s.Pop()
	s.stack1 = append(s.stack1, temp)
	return temp
}

func (s *CosplayQueue) Pop() (res int) {
	if len(s.stack2) == 0 {
		for i := len(s.stack1) - 1; i >= 0; i-- {
			s.stack2 = append(s.stack2, s.stack1[i])
			s.stack1 = s.stack1[:i]
		}
	}

	res = s.stack2[len(s.stack2)-1]
	s.stack2 = s.stack2[:len(s.stack2)-1]
	return
}

func (s *CosplayQueue) Empty() bool {
	return len(s.stack1) == 0 && len(s.stack2) == 0
}
