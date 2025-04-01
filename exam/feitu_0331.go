package exam

import "fmt"

type StructA struct {
	Times []int
	// check_type string
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}

// pre time next

func (s *StructA) getClosest(time, res int) int {
	if res == 0 {
		return s.Times[res]
	}
	if res > len(s.Times)-1 {
		return s.Times[len(s.Times)-1]
	}
	pre, next := res-1, res
	if time-s.Times[pre] > s.Times[next]-time {
		return s.Times[next]
	}
	return s.Times[pre]
}

func (s *StructA) getEarlier(time, res int) int {
	if res == 0 {
		return s.Times[res]
	}
	return s.Times[res-1]
}

func (s *StructA) getNext(time, res int) int {
	if res > len(s.Times)-1 {
		return s.Times[res-1]
	}
	return s.Times[res]
}

func (s *StructA) GetByTime(time int) {
	res := binarySearch(s.Times, time)
	if 0 <= res && res <= len(s.Times)-1 && s.Times[res] == time {
		fmt.Printf("精确匹配到%d\n", s.Times[res])
		return
	}
	res1, res2, res3 := s.getClosest(time, res), s.getEarlier(time, res), s.getNext(time, res)
	fmt.Printf("优先最近结果是%d, 优先更早结果是%d, 优先更晚结果是%d\n", res1, res2, res3)
}

// var getInfo interface {
// 	func()
// }
