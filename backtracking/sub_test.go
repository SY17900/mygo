package backtracking

import (
	"fmt"
	"testing"
)

func TestSubSet1(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := SubSet1(nums)
	fmt.Println(ans)
}

func TestSubSet2(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	ans := SubSet2(nums)
	fmt.Println(ans)
}

func TestIncreasingSSQ(t *testing.T) {
	nums := []int{4, 6, 6, 7}
	ans := IncreasingSSQ(nums)
	fmt.Println(ans)
}
