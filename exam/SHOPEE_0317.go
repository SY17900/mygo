package exam

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SHOPEE_EXAM1() {
	var str string
	fmt.Scanln(&str)

	startIndex, ans := 0, 1
	rec := make(map[byte]int)
	for i := range str {
		if dIndex, ok := rec[str[i]]; ok {
			ans = max(ans, i-startIndex)
			for j := startIndex; j <= dIndex; j++ {
				delete(rec, str[j])
			}
			rec[str[i]] = i
			startIndex = dIndex + 1
			continue
		}
		rec[str[i]] = i
	}
	ans = max(ans, len(str)-startIndex)

	fmt.Println(ans)
}

func SHOPEE_EXAM2() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	temp := sc.Text()

	num1 := make([]int, 0)
	num2 := make([]int, 0)
	fs, fe := strings.IndexByte(temp, '['), strings.IndexByte(temp, ']')
	ss, se := strings.LastIndexByte(temp, '['), strings.LastIndexByte(temp, ']')
	tempS1 := strings.Split(temp[fs+1:fe], ",")
	tempS2 := strings.Split(temp[ss+1:se], ",")
	for _, str := range tempS1 {
		num, _ := strconv.Atoi(str)
		num1 = append(num1, num)
	}
	for _, str := range tempS2 {
		num, _ := strconv.Atoi(str)
		num2 = append(num2, num)
	}

	ans := make([]int, len(num1)+len(num2))
	index := len(ans) - 1
	var i, j int
	for i, j = len(num1)-1, len(num2)-1; i >= 0 && j >= 0; {
		if num1[i] >= num2[j] {
			ans[index] = num1[i]
			i--
		} else {
			ans[index] = num2[j]
			j--
		}
		index--
	}

	if i >= 0 {
		for m := 0; m <= i; m++ {
			ans[m] = num1[m]
		}
	}
	if j >= 0 {
		for m := 0; m <= j; m++ {
			ans[m] = num2[m]
		}
	}

	fmt.Print("[")
	for i := range ans {
		fmt.Print(ans[i])
		if i != len(ans)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}

func SHOPEE_EXAM3() {
	var str string
	var k int
	fmt.Scan(&str, &k)

	ifValid := func(chr byte) bool {
		if chr == 'a' || chr == 'e' || chr == 'i' || chr == 'o' || chr == 'u' {
			return true
		}
		return false
	}

	ans, rec := 0, 0
	for i := range k {
		if ifValid(str[i]) {
			rec++
		}
	}
	ans = rec
	if ans == k {
		fmt.Println(ans)
		return
	}

	for i := k; i < len(str); i++ {
		if ifValid(str[i-k]) {
			rec--
		}
		if ifValid(str[i]) {
			rec++
		}
		ans = max(ans, rec)
		if ans == k {
			fmt.Println(ans)
			return
		}
	}

	fmt.Println(ans)
}
