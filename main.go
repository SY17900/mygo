package main

import (
	"fmt"
	"mygo/sorting"
)

func main() {
	nums := []int{5, 2, 4, 6, 1, 3}
	sorting.HeapSort(nums)
	for index, num := range nums {
		fmt.Print(num)
		if index != len(nums)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
