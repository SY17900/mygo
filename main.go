package main

import "fmt"

func main() {
	slice := []byte{'n', 'a', 'n', 'o'}
	add := func() {
		slice = append(slice, 'd')
	}
	add()
	fmt.Println(slice)
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
	add()
	fmt.Println(slice)
}
