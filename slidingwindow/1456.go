package slidingwindow

func ifVowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func MaxVowels(s string, k int) (ans int) {
	cnt := 0
	for i, in := range s {
		if ifVowel(byte(in)) {
			cnt++
		}
		if i < k-1 {
			continue
		}

		ans = max(ans, cnt)
		out := s[i-k+1]
		if ifVowel(out) {
			cnt--
		}
	}

	return
}
