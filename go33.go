package main

import "fmt"

func main() {
	h := "hello world"

	for idx, c := range h {
		fmt.Printf("%d %c %c\n", idx, h[idx], c)
	}
}
