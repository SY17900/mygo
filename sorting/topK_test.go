package sorting

import (
	"testing"
)

func TestTopK(t *testing.T) {
	topK := NewTopK(3)
	topK.Push(5)
	topK.Push(3)
	topK.Push(8)
	topK.Push(2)
	topK.Push(9)
	topK.Push(11)

	topK.PrintOut()
}
