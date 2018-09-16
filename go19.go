package main

import "fmt"

func removeElem(stack []int, idx int) (int, []int) {
	middle := stack[idx]
	copy(stack[idx:], stack[idx+1:])
	return middle, stack[:len(stack)-1]
}

func main() {
	x := make([]int, 0)

	for i := 0; i < 9; i++ {
		x = append(x, i)
	}

	middle, res := removeElem(x, 4)
	fmt.Println(middle, res)
}
