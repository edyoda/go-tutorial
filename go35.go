package main

import (
	"fmt"
)

//func sum(z []int)

func main() {
	a := make([]string, 0)
	b := make([]int, 5)

	c := append(a, "hello", "world", "here")

	fmt.Println(c)
	fmt.Println(a)

	z := append(b, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	sum := 0
	for _, v := range z {
		sum += v
	}

	println(sum)

}
