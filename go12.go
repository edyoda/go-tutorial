package main

import (
	"fmt"
)

func main() {
	s := "hello world"

	var n int
	for _, _ = range s {
		n++
	}

	fmt.Println(s[2:5], n)
}
