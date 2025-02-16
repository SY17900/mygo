// 是否是有效的ipv4地址
package others

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ifvalid(num string) bool {
	if len(num) <= 0 || len(num) > 3 {
		return false
	}
	if num[0] == '0' && len(num) != 1 {
		return false
	}
	i, err := strconv.Atoi(num)
	if err != nil {
		return false
	}

	return 0 <= i && i <= 255
}

func IPaddr() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	parts := strings.Split(text, ".")
	if len(parts) != 4 {
		fmt.Println("Invalid")
		return
	}

	for _, num := range parts {
		if !ifvalid(num) {
			fmt.Println("Invalid")
			return
		}
	}

	fmt.Println("Valid")
}
