package strings

func makeNext(arr string) []int {
	next := make([]int, len(arr))
	next[0] = 0
	j := 0
	for i := 1; i < len(arr); i++ {
		for j > 0 && arr[j] != arr[i] {
			j = next[j-1]
		}
		if arr[j] == arr[i] {
			j++
		}
		next[i] = j
	}

	return next
}

func StrStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := makeNext(needle)
	j := 0
	for i := range len(haystack) {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}
