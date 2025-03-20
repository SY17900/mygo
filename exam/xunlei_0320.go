package exam

import (
	"fmt"
	"strconv"
)

func XUNLEI_EXAM1() {
	var n int
	fmt.Scan(&n)
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	num := 1
	for startRow := 0; startRow < n/2; startRow++ {
		steps := n - startRow*2 - 1
		for move := 0; move < steps; move++ {
			ans[startRow][startRow+move] = num
			num++
		}
		for move := 0; move < steps; move++ {
			ans[startRow+move][n-1-startRow] = num
			num++
		}
		for move := 0; move < steps; move++ {
			ans[n-1-startRow][n-1-startRow-move] = num
			num++
		}
		for move := 0; move < steps; move++ {
			ans[n-1-startRow-move][startRow] = num
			num++
		}
	}
	if n%2 != 0 {
		ans[n/2][n/2] = num
	}

	for i := range ans {
		fmt.Println(ans[i])
	}
}

func XUNLEI_EXAM2() {
	var num1, num2 string
	fmt.Scan(&num1, &num2)
	ifNegative := false
	if num1[0] == '-' && num2[0] == '-' {
		num1 = num1[1:]
		num2 = num2[1:]
	}
	if num1[0] == '-' {
		num1 = num1[1:]
		ifNegative = true
	} else if num2[0] == '-' {
		num2 = num2[1:]
		ifNegative = true
	}

	ans := helper(num1, num2)

	if ifNegative {
		ans = "-" + ans
	}
	fmt.Println(ans)
}

func helper(str1, str2 string) string {
	len2 := len(str2)
	var ans string
	for i := 0; i < len2; i++ {
		num2 := int(str2[len2-1-i] - '0')
		tenTimes := i
		temp := str1
		for addTimes := 0; addTimes < num2; addTimes++ {
			temp = Phelper(temp, str1)
		}
		for plusTimes := 0; plusTimes < tenTimes; plusTimes++ {
			temp = temp + "0"
		}
		ans = Phelper(ans, temp)
	}

	return ans
}

func Phelper(str1, str2 string) string {
	pos1, pos2, pre, this := len(str1)-1, len(str2)-1, 0, 0
	var ans string
	for ; pos1 >= 0 || pos2 >= 0; pos1, pos2 = pos1-1, pos2-1 {
		var temp1, temp2 byte
		var num1, num2 int
		if pos1 >= 0 {
			temp1 = str1[pos1]
			num1 = int(temp1 - '0')
		}
		if pos2 >= 0 {
			temp2 = str2[pos2]
			num2 = int(temp2 - '0')
		}
		this, pre = (num1+num2+pre)%10, (num1+num2+pre)/10
		ans = strconv.Itoa(this) + ans
	}
	if pre != 0 {
		ans = strconv.Itoa(pre) + ans
	}

	return ans
}
