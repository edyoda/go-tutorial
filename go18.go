package main

import (
	"fmt"
)

func main() {
	stack := make([]int, 0)

	//push 5
	stack = append(stack, 5)

	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	fmt.Println(top)
}
