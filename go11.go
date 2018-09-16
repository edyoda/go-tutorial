package main

import "fmt"

type Far float32

func main() {
	i := 10
	j := 20

	var s Far

	i, j = j, i

	s = 1.2
	fmt.Println(i, j, s)
}
