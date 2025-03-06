package strings

func Reverse(str []byte) {
	for left, right := 0, len(str)-1; left < right; {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}

func Reverse2k(str []byte, k int) {
	for i := 0; i < len(str); i += 2 * k {
		if i+k < len(str) {
			Reverse(str[i : i+k])
		} else {
			Reverse(str[i:])
		}
	}
}

func ReverseWords(str []byte) []byte {
	i := 0
	for ; str[i] == ' '; i++ {
	}
	str = str[i:]
	i = len(str) - 1
	for ; str[i] == ' '; i-- {
	}
	str = str[:i+1]

	var slow, fast int
	for ; fast < len(str); fast++ {
		if str[fast] == ' ' && str[fast] == str[fast-1] {
			continue
		}
		str[slow] = str[fast]
		slow++
	}

	str = str[:slow]
	Reverse(str)
	slow, fast = 0, 0
	for fast < len(str) {
		for ; fast < len(str) && str[fast] != ' '; fast++ {

		}
		Reverse(str[slow:fast])
		fast++
		slow = fast
	}

	return str
}

func MoveRight(str []byte, n int) {
	Reverse(str)
	Reverse(str[:n])
	Reverse(str[n:])
}
