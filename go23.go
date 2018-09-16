package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 20

	fmt.Println(add(a, b))

	a = 30

	fmt.Println(add(a, b))

}
