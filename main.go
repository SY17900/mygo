package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	for ; n > 0; n-- {
		sc.Scan()
		line := sc.Text()
		words := strings.Fields(line)
		ans := make([]byte, 0, len(words))
		for _, word := range words {
			ans = append(ans, word[0])
		}
		fmt.Println([]byte(strings.ToUpper(string(ans))))
	}
}
