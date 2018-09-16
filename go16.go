package main

import (
	"fmt"
)

func appendMore(x []int, y []int) []int {

	zlen := len(x) + len(y)
	z := make([]int, zlen)
	copy(z, x)
	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int

	for i := 0; i < 10; i++ {
		x = append(x, i)
		y = append(y, i)
		// fmt.Println(cap(x,10))
	}

	fmt.Println(x)
	ZZ := appendMore(x, y)
	fmt.Println(ZZ)

}
