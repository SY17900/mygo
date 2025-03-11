package backtracking

import (
	"fmt"
	"testing"
)

func TestFullPermutationD(t *testing.T) {
	nums := []int{1, 1, 2}
	ans := FullPermutationD1(nums)
	fmt.Println(ans)
}
