package backtracking

import (
	"fmt"
	"testing"
)

func TestCombineK(t *testing.T) {
	ans := CombineK(4, 2)
	fmt.Println(ans)
}

func TestCombineSum(t *testing.T) {
	ans := CombineSum(9, 3)
	fmt.Println(ans)
}

func TestLetters(t *testing.T) {
	ans := Letters("23")
	fmt.Println(ans)
}

func TestCombineSumD(t *testing.T) {
	nums := []int{2, 3, 5}
	ans := CombineSumD(nums, 8)
	fmt.Println(ans)
}

func TestCombineSumND(t *testing.T) {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	ans := CombineSumND(nums, 8)
	fmt.Println(ans)
}

func TestMakePalindrome(t *testing.T) {
	ans := MakePalindrome("aab")
	fmt.Println(ans)
}
