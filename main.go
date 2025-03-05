package main

import (
	"fmt"
	"mygo/strings"
)

func main() {
	str := []byte("abcdefg")
	strings.MoveRight(str, 2)
	var res string = string(str)
	fmt.Println(res)
}
