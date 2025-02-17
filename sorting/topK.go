package sorting

type TopK struct {
	size int
	data []int
}

func NewTopK(size int) *TopK {
	return &TopK{
		size: size,
		data: make([]int, size),
	}
}

func (t *TopK) down(index int) {
	parent := index
	child := index * 2
	for child < t.size {
		if child+1 < t.size && t.data[child+1] < t.data[child] {
			child++
		}
		if t.data[parent] <= t.data[child] {
			break
		}
		t.data[parent], t.data[child] = t.data[child], t.data[parent]
		parent = child
		child = parent * 2
	}
}

func (t *TopK) Push(val int) {
	if len(t.data) == t.size {
		if val <= t.data[0] {
			return
		}
		t.data[0] = val
		t.down(0)
	} else {
		t.data = append(t.data, val)
		if len(t.data) == t.size {
			for i := len(t.data)/2 - 1; i >= 0; i-- {
				t.down(i)
			}
		}
	}
}

func (t *TopK) PrintOut() {
	for index, num := range t.data {
		print(num)
		if index != t.size-1 {
			print(" ")
		}
	}
	println()
}
