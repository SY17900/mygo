package hashmap

func Word(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	rec := [26]int{}
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		sPosition, tPosition := s[i]-'a', t[i]-'a'
		rec[sPosition]++
		rec[tPosition]--
	}

	for _, cnt := range rec {
		if cnt != 0 {
			return false
		}
	}
	return true
}

func Intersection(nums1, nums2 []int) []int {
	set := make(map[int]struct{})
	res := make([]int, 0, 26)

	for _, num := range nums1 {
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
		}
	}
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			res = append(res, num)
			delete(set, num)
		}
	}

	return res
}

func IsHappy(num int) bool {
	rec := make(map[int]struct{})
	for num != 1 {
		if _, ok := rec[isHaapyHelper(num)]; ok {
			return false
		}
		rec[num] = struct{}{}
	}
	return true
}

func isHaapyHelper(num int) (sum int) {
	for num != 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return
}
