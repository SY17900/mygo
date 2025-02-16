package course

// 判断两个字符串是否相互是打乱顺序的
func IfAnagrams(s1, s2 string) bool {
	if len(s1) != len((s2)) {
		return false
	}

	mp := make(map[rune]int)

	for _, ch := range s1 {
		mp[ch]++
	}

	for _, ch := range s2 {
		mp[ch]--
	}

	for _, count := range mp {
		if count != 0 {
			return false
		}
	}

	return true
}
