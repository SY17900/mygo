package backtracking

import (
	"fmt"
	"testing"
)

func TestNQueen(t *testing.T) {
	ans := NQueens(4)
	fmt.Println(ans)
}
